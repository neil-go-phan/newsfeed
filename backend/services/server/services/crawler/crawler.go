package crawlerservices

import (
	"fmt"
	"net/url"
	"os"
	"server/entities"
	"server/helpers"
	pb "server/proto"
	"server/repository"
	"server/services"
	"strings"

	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

type CrawlerService struct {
	repo                   repository.CrawlerRepository
	articleService         services.ArticleServices
	articlesSourceServices services.ArticlesSourceServices
	cronjobService         services.CronjobServices
	grpcClient             pb.CrawlerServiceClient
	roleServices           services.RoleServices
}

const DEFAULT_SCHEDULE = "@every 0h5m"
const CRAWLER_ROLE_ENTITY = "CRAWLER"
const CRAWLER_ROLE_CREATE_METHOD = "CREATE"
const CRAWLER_ROLE_UPDATE_METHOD = "UPDATE"


func NewCrawlerService(repo repository.CrawlerRepository, articleService services.ArticleServices, articlesSourceServices services.ArticlesSourceServices, cronjobService services.CronjobServices, grpcClient pb.CrawlerServiceClient, roleServices services.RoleServices) *CrawlerService {
	crawlerService := &CrawlerService{
		repo:                   repo,
		articleService:         articleService,
		articlesSourceServices: articlesSourceServices,
		cronjobService:         cronjobService,
		grpcClient:             grpcClient,
		roleServices:           roleServices,
	}
	return crawlerService
}

func (s *CrawlerService) TestRSSCrawler(crawler entities.Crawler) (*services.ArticlesSourceResponseCrawl, []*services.ArticleResponse, error) {
	err := validateCrawler(crawler)
	if err != nil {
		return nil, nil, err
	}

	result, err := getTestRSSCrawlerResult(s.grpcClient, crawler)
	if err != nil {
		return nil, nil, err
	}

	articlesSource, articles := helpers.CastTestResult(result)
	return articlesSource, articles, nil
}

func (s *CrawlerService) TestCustomCrawler(crawler entities.Crawler) (*services.ArticlesSourceResponseCrawl, []*services.ArticleResponse, error) {
	err := validateCrawler(crawler)
	if err != nil {
		return nil, nil, err
	}
	result, err := getTestCustomCrawlerResult(s.grpcClient, crawler)

	if err != nil {
		return nil, nil, err
	}

	articlesSource, articles := helpers.CastTestResult(result)
	return articlesSource, articles, nil
}

func (s *CrawlerService) CreateCrawlerWithCorrespondingArticlesSource(role string, payload services.CreateCrawlerPayload) error {
	isAllowed := s.roleServices.GrantPermission(role, CRAWLER_ROLE_ENTITY, CRAWLER_ROLE_CREATE_METHOD)
	if !isAllowed {
		return fmt.Errorf("unauthorized")
	}

	articlesSource, crawler := extractPayload(payload)
	err := validateCreateCrawlerPayload(articlesSource, crawler)
	if err != nil {
		return fmt.Errorf("validate error: %s", err.Error())
	}

	articlesSource, err = s.articlesSourceServices.CreateIfNotExist(articlesSource)
	if err != nil {
		return err
	}

	crawler.Schedule = DEFAULT_SCHEDULE
	crawler.ArticlesSourceID = articlesSource.ID

	crawler, err = s.repo.CreateIfNotExist(crawler)
	if err != nil {
		return err
	}

	s.cronjobService.CreateCrawlerCronjob(crawler)

	return nil
}

func (s *CrawlerService) GetHtmlPage(url *url.URL) error {
	ctx, cancel := createContext()
	defer cancel()
	var htmlContent string
	task := getDocTask(url.String(), &htmlContent)

	if err := chromedp.Run(ctx, task); err != nil {
		fmt.Println(err)
	}

	hostname := strings.TrimPrefix(url.Hostname(), "www.")
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		log.Error(err)
	}

	removeScriptTags(doc)

	htmlWithoutScript := renderNode(doc)
	err = os.WriteFile(fmt.Sprintf("page%s.html", hostname), []byte(htmlWithoutScript), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (s *CrawlerService) CreateCrawlerCronjobFromDB() error {
	crawlers, err := s.repo.List()
	if err != nil {
		return err
	}
	for _, crawler := range crawlers {
		s.cronjobService.CreateCrawlerCronjob(crawler)
	}
	return nil
}

func (s *CrawlerService) ListAllPaging(page int, pageSize int) ([]services.CrawlerResponse, int64, error) {
	crawlerResponse := make([]services.CrawlerResponse, 0)
	crawlers, found, err := s.repo.ListAllPaging(page, pageSize)
	if err != nil {
		return crawlerResponse, found, err
	}
	for _, crawler := range crawlers {
		crawlerResponse = append(crawlerResponse, castCrawlerToResponse(crawler))
	}
	return crawlerResponse, found, nil
}

func (s *CrawlerService) UpdateSchedule(role string, id uint, newSchedule string) error {
	isAllowed := s.roleServices.GrantPermission(role, CRAWLER_ROLE_ENTITY, CRAWLER_ROLE_UPDATE_METHOD)
	if !isAllowed {
		return fmt.Errorf("unauthorized")
	}
	crawler, err := s.repo.Get(id)
	if err != nil {
		return err
	}

	err = s.cronjobService.RemoveCronjob(*crawler)
	if err != nil {
		return err
	}

	crawler.Schedule = newSchedule

	s.cronjobService.CreateCrawlerCronjob(*crawler)

	err = s.repo.UpdateSchedule(id, newSchedule)
	if err != nil {
		return err
	}

	return nil
}

func (s *CrawlerService) Get(id uint) (*entities.Crawler, error) {
	return s.repo.Get(id)
}

func (s *CrawlerService) Update(role string, crawler entities.Crawler) error {
	isAllowed := s.roleServices.GrantPermission(role, CRAWLER_ROLE_ENTITY, CRAWLER_ROLE_UPDATE_METHOD)
	if !isAllowed {
		return fmt.Errorf("unauthorized")
	}
	return s.repo.Update(crawler)
}

func (s *CrawlerService) CronjobOnDay(timeString string) (*[24]services.ChartDay, error) {
	return s.cronjobService.CronjobOnDay(timeString)
}

func (s *CrawlerService) CronjobOnHour(timeString string) (*[60]services.ChartHour, error) {
	return s.cronjobService.CronjobOnHour(timeString)
}

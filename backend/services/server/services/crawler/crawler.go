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
	grpcClient             pb.CrawlerServiceClient
}

func NewCrawlerService(repo repository.CrawlerRepository, articleService services.ArticleServices, articlesSourceServices services.ArticlesSourceServices, grpcClient pb.CrawlerServiceClient) *CrawlerService {
	crawlerService := &CrawlerService{
		repo:                   repo,
		articleService:         articleService,
		articlesSourceServices: articlesSourceServices,
		grpcClient:             grpcClient,
	}
	return crawlerService
}

func (s *CrawlerService) TestRSSCrawler(crawler entities.Crawler) (*services.ArticlesSourceResponse, []*services.ArticleResponse, error) {
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

func (s *CrawlerService) TestCustomCrawler(crawler entities.Crawler)  (*services.ArticlesSourceResponse, []*services.ArticleResponse, error) {
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
func (s *CrawlerService) CreateCrawlerWithCorrespondingArticlesSource(payload services.CreateCrawlerPayload) error {
	articlesSource, crawler := exportPayload(payload)
	err := validateCreateCrawlerPayload(articlesSource, crawler)
	if err != nil {
		return fmt.Errorf("validate error: %s", err.Error())
	}

	err = s.articlesSourceServices.CreateIfNotExist(articlesSource)
	if err != nil {
		return err
	}

	err = s.repo.CreateIfNotExist(crawler)
	if err != nil {
		return err
	}
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

func (s *CrawlerService) FirstCrawl(crawler *entities.Crawler) error {
	return nil
}

func (s *CrawlerService) ScheduledCrawl(crawlerID uint) error {
	return nil
}

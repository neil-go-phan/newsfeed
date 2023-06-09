package crawlerservices

import (
	"server/entities"
	"server/helpers"
	pb "server/proto"
	"server/repository"
	"server/services"
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

func (s *CrawlerService) TestCrawler(crawler *entities.Crawler) (*entities.ArticlesSource, []*entities.Article, error) {
	err := validateCrawler(crawler)
	if err != nil {
		return nil, nil, err
	}
	
	result, err := getTestCrawlerResult(s.grpcClient, crawler)
	if err != nil {
		return nil, nil, err
	}
	
	articlesSource, articles := helpers.CastTestResult(result)
	return articlesSource, articles, nil
}

func (s *CrawlerService) FirstCrawl(crawler *entities.Crawler) error {
	return nil
}

func (s *CrawlerService) ScheduledCrawl(crawlerID uint) error {
	return nil
}

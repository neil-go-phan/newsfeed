package crawlerservices

import (
	"crawler/entities"
	"crawler/repository"
	"crawler/services"
)

type CrawlerService struct {
	repo                   repository.CrawlerRepository
	articleService         services.ArticleServices
	articlesSourceServices services.ArticlesSourceServices
}

func NewCrawlerService(repo repository.CrawlerRepository, articleService services.ArticleServices, articlesSourceServices services.ArticlesSourceServices) *CrawlerService {
	crawlerService := &CrawlerService{
		repo: repo,
		articleService: articleService,
		articlesSourceServices:articlesSourceServices,
	}
	return crawlerService
}

func (s *CrawlerService) TestCrawler(crawler entities.Crawler) (*entities.ArticlesSource, []*entities.Article, error) {
	articleSource, articles, err := TestCrawlWithRSS(crawler)
	if err != nil {
		return nil, nil, err
	}

	return articleSource, articles, nil
}



func (s *CrawlerService) FirstCrawl(crawler entities.Crawler) (error) {
	return nil
}

func (s *CrawlerService) ScheduledCrawl(crawlerID uint) (error) {
	return nil
}
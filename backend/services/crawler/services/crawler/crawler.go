package crawlerservices

import (
	"crawler/entities"
	"crawler/repository"
	"crawler/services"
	log "github.com/sirupsen/logrus"
)

const CRAWLER_FEED_TYPE = "feed"
const CRAWLER_CUSTOM_TYPE = "custom"

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

func (s *CrawlerService) TestRSSCrawler(crawler entities.Crawler) (entities.ArticlesSource, []entities.Article, error) {
	articleSource, articles, err := TestCrawlWithRSS(crawler)
	if err != nil {
		return articleSource, articles, err
	}

	return articleSource, articles, nil
}

func (s *CrawlerService) TestCustomCrawler(crawler entities.Crawler) (entities.ArticlesSource, []entities.Article, error) {
	articleSource, articles, err := TestCustomCrawl(crawler)
	if err != nil {
		return articleSource, articles, err
	}

	return articleSource, articles, nil
}

func (s *CrawlerService) Crawl(crawler entities.Crawler) (newArticleCount int32, err error) {
	articles := make([]entities.Article, 0)
	if crawler.CrawlType == CRAWLER_FEED_TYPE {
		articles, err = CrawlWithRSS(crawler)
		if err != nil {
			return newArticleCount, err
		}
	}

	if crawler.CrawlType == CRAWLER_CUSTOM_TYPE {
		articles, err = CrawlWithCustomCrawler(crawler)
		if err != nil {
			return newArticleCount, err
		}
	}

	newArticleCount = s.articleService.StoreArticles(articles, crawler.ArticlesSourceID)
	log.Printf("Found %v new articles", newArticleCount)

	return newArticleCount, nil
}

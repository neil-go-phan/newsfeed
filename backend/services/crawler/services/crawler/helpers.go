package crawlerservices

import (
	"crawler/crawl"
	"crawler/entities"
	"crawler/helpers"
)

func TestCrawlWithRSS(crawler entities.Crawler) (*entities.ArticlesSource, []*entities.Article, error) {
	feed, err := crawl.TestGetRSSFeed(crawler)
	if err != nil {
		return nil, nil, err
	}

	articleSource := helpers.CastFeedToArticlesSource(feed)
	articles := helpers.CastFeedItemsToArticles(feed.Items)

	return articleSource, articles, nil
}

func TestCustomCrawl(crawler entities.Crawler) (entities.ArticlesSource, []*entities.Article, error) {
	articles, err := crawl.CrawlWithGoQuery(crawler)
	if err != nil {
		return entities.ArticlesSource{}, []*entities.Article{}, err
	}

	articleSource := entities.ArticlesSource{
		Link: crawler.SourceLink,
	}

	return articleSource, articles, nil
}

func TestCrawlWithGoQuery(crawler *entities.Crawler) {

}
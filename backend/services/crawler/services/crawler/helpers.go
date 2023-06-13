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



func TestCrawlWithGoQuery(crawler *entities.Crawler) {

}
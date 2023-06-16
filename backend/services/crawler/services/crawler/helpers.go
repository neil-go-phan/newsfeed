package crawlerservices

import (
	"crawler/crawl"
	"crawler/entities"
	"crawler/helpers"
	log "github.com/sirupsen/logrus"
)

func TestCrawlWithRSS(crawler entities.Crawler) (entities.ArticlesSource, []entities.Article, error) {
	feed, err := crawl.TestGetRSSFeed(crawler)
	if err != nil {
		return entities.ArticlesSource{}, []entities.Article{}, err
	}

	articleSource := helpers.CastFeedToArticlesSource(feed)
	articles := helpers.CastFeedItemsToArticles(feed.Items)

	return articleSource, articles, nil
}

func TestCustomCrawl(crawler entities.Crawler) (entities.ArticlesSource, []entities.Article, error) {
	articles, err := crawl.CrawlWithGoQuery(crawler)
	if err != nil {
		return entities.ArticlesSource{}, []entities.Article{}, err
	}

	if helpers.CheckEmptyArticles(articles) {
		log.Println("not found article try crawl with chromedp")
		articles, err = crawl.CrawlWithChromedp(crawler)
		if err != nil {
			return entities.ArticlesSource{}, []entities.Article{}, err
		}
	}

	articleSource := entities.ArticlesSource{
		Link: crawler.SourceLink,
	}

	return articleSource, articles, nil
}

func CrawlWithRSS(crawler entities.Crawler) ([]entities.Article, error){
	feed, err := crawl.GetRSSFeed(crawler.FeedLink)
	if err != nil {
		return nil, err
	}
	articles := helpers.CastFeedItemsToArticles(feed.Items)
	return articles, nil
}

func CrawlWithCustomCrawler(crawler entities.Crawler) ([]entities.Article, error) {
	articles, err := crawl.CrawlWithGoQuery(crawler)
	if err != nil {
		return []entities.Article{}, err
	}

	if helpers.CheckEmptyArticles(articles) {
		log.Println("not found article try crawl with chromedp")
		articles, err = crawl.CrawlWithChromedp(crawler)
		if err != nil {
			return []entities.Article{}, err
		}
	}

	return articles, nil
}

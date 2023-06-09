package crawlerservices

import (
	"crawler/crawl"
	"crawler/entities"

	"github.com/mmcdole/gofeed"
)

func TestCrawlWithRSS(crawler *entities.Crawler) (*entities.ArticlesSource, []*entities.Article, error) {
	feed, err := crawl.TestGetRSSFeed(crawler)
	if err != nil {
		return nil, nil, err
	}

	articleSource, articles := castFeedToResponse(feed)

	return articleSource, articles, nil
}

func castFeedToResponse(feed *gofeed.Feed) (*entities.ArticlesSource, []*entities.Article) {
	articleSource := castFeedToArticlesSource(feed)
	articles := castFeedItemsToArticles(feed.Items)
	return articleSource, articles
}

func castFeedToArticlesSource(feed *gofeed.Feed) *entities.ArticlesSource {
	return &entities.ArticlesSource{
		Title: feed.Title,
		Description: feed.Description,
		Link: feed.Link,
		FeedLink: feed.FeedLink,
		Image: feed.Image.URL,
	}
}

func castFeedItemsToArticles(items []*gofeed.Item) []*entities.Article {
	articles := make([]*entities.Article, 0)
	for _, item := range items {
		articles = append(articles, castFeedItemToArticle(item))
	}
	return articles
}

func castFeedItemToArticle(item *gofeed.Item) (*entities.Article){
	return &entities.Article{
		Title: item.Title,
		Description: item.Description,
		Link: item.Link,
		Published: *item.PublishedParsed,
		Authors: createAuthorsNameString(item.Authors),
	}
}

func createAuthorsNameString(authors []*gofeed.Person) (string) {
	authorsNameString := ""
	for _, author := range authors {
		authorsNameString += author.Name 
		authorsNameString += " "
	}
	return authorsNameString
}

func TestCrawlWithGoQuery(crawler *entities.Crawler) {

}
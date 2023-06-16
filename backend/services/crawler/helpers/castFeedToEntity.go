package helpers

import (
	"crawler/entities"

	"github.com/mmcdole/gofeed"
)

func CastFeedToArticlesSource(feed *gofeed.Feed) entities.ArticlesSource {
	articleSource := entities.ArticlesSource{
		Title: feed.Title,
		Description: feed.Description,
		Link: feed.Link,
		FeedLink: feed.FeedLink,
	}
	if (feed.Image != nil) {
		articleSource.Image = feed.Image.URL
	}
	return articleSource
}

func CastFeedItemsToArticles(items []*gofeed.Item) []entities.Article {
	articles := make([]entities.Article, 0)
	for _, item := range items {
		articles = append(articles, castFeedItemToArticle(item))
	}
	return articles
}

func castFeedItemToArticle(item *gofeed.Item) (entities.Article){
	article := entities.Article{
		Title: item.Title,
		Description: item.Description,
		Link: item.Link,
	}
	if (item.PublishedParsed != nil) {
		article.Published = *item.PublishedParsed
	}
	if (len(item.Authors) != 0) {
		article.Authors = createAuthorsNameString(item.Authors)
	}
	return article
}

func createAuthorsNameString(authors []*gofeed.Person) (string) {
	authorsNameString := ""
	for _, author := range authors {
		authorsNameString += author.Name 
		authorsNameString += " "
	}
	return authorsNameString
}

// func CastEntityArticleToPbArticles(items []*gofeed.Item) []*entities.Article {
// 	articles := make([]*entities.Article, 0)
// 	for _, item := range items {
// 		articles = append(articles, castFeedItemToArticle(item))
// 	}
// 	return articles
// }
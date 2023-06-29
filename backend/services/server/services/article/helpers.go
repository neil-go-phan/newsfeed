package articleservices

import (
	"server/entities"
	"server/repository"
	"server/services"
)

func castEntityArticleToReponse(article entities.Article) services.ArticleResponse {
	return services.ArticleResponse{
		ID: article.ID,
		Title: article.Title,
		Description: article.Description,
		Link: article.Link,
		Published: article.Published,
		Authors: article.Authors,
		ArticlesSourceID: article.ArticlesSourceID,
	}
}

func castArticleFromRepoToArticleReadReponse(article repository.ArticleLeftJoinRead) services.ArticleForReadResponse {
	newArticle := services.ArticleForReadResponse{
		ID: article.ID,
		Title: article.Title,
		Description: article.Description,
		Link: article.Link,
		Published: article.Published,
		Authors: article.Authors,
		ArticlesSourceID: article.ArticlesSourceID,
		IsRead: article.IsRead,
		IsReadLater: article.IsReadLater,
	}
	return newArticle
}
package articleservices

import (
	"server/entities"
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
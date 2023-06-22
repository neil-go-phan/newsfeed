package articlessourceservices

import (
	"server/entities"
	"server/services"
)

func castEntityArticlesSourceToReponse(articlesSource entities.ArticlesSource) services.ArticlesSourceResponseRender {
	return services.ArticlesSourceResponseRender{
		ID: articlesSource.ID,
		Title: articlesSource.Title,
		Description: articlesSource.Description,
		Link: articlesSource.Link,
		Image: articlesSource.Image,
		TopicID: articlesSource.TopicID,
	}
}
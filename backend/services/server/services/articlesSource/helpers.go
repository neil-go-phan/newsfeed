package articlessourceservices

import (
	"server/entities"
	"server/repository"
	"server/services"
)

func castEntityArticlesSourceToReponse(articlesSource entities.ArticlesSource) services.ArticlesSourceResponseRender {
	return services.ArticlesSourceResponseRender{
		ID: articlesSource.ID,
		Title: articlesSource.Title,
		Description: articlesSource.Description,
		Link: articlesSource.Link,
		Image: articlesSource.Image,
		Follower: articlesSource.Follower,
		TopicID: articlesSource.TopicID,
	}
}

func castArticlesSourceRecommended(articlesSource repository.MostActiveSource) services.ArticlesSourceRecommended {
	return services.ArticlesSourceRecommended{
		ID: articlesSource.ID,
		Title: articlesSource.Title,
		Description: articlesSource.Description,
		Link: articlesSource.Link,
		Image: articlesSource.Image,
		Follower: articlesSource.Follower,
		TopicID: articlesSource.TopicID,
		ArticlesPreviousWeek: articlesSource.ArticlesPreviousWeek,
	}
}

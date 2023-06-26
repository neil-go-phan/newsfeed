package followservices

import (
	"server/entities"
	"server/services"
)

func getArticlesSourcesFromArrayFollow(follows []entities.Follow) []services.ArticlesSourceResponseRender {
	articlesSources := make([]services.ArticlesSourceResponseRender, 0)
	for _,follow := range follows {
		serviceSource := services.ArticlesSourceResponseRender{
			ID: follow.ArticlesSource.ID,
			Title: follow.ArticlesSource.Title,
			Description: follow.ArticlesSource.Description,
			Link: follow.ArticlesSource.Link,
			Image: follow.ArticlesSource.Image,
			Follower: follow.ArticlesSource.Follower,
			TopicID: follow.ArticlesSource.TopicID,
		}
		articlesSources = append(articlesSources, serviceSource)
	}
	return articlesSources;
}


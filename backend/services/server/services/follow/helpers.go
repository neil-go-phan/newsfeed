package followservices

import (
	"server/entities"
	"server/services"
)

func getArticlesSourcesFromArrayFollow(follows []entities.Follow) []services.ArticlesSourceUserFollow {
	articlesSources := make([]services.ArticlesSourceUserFollow, 0)
	for _,follow := range follows {
		serviceSource := services.ArticlesSourceUserFollow{
			ID: follow.ArticlesSource.ID,
			Title: follow.ArticlesSource.Title,
			Description: follow.ArticlesSource.Description,
			Link: follow.ArticlesSource.Link,
			Image: follow.ArticlesSource.Image,
			Follower: follow.ArticlesSource.Follower,
			TopicID: follow.ArticlesSource.TopicID,
			Unread: follow.Unread,
		}
		articlesSources = append(articlesSources, serviceSource)
	}
	return articlesSources;
}


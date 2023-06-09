package helpers

import (
	"crawler/entities"
	pb "crawler/proto"
	log "github.com/sirupsen/logrus"
	"encoding/json"
)

func CastEntityArticleSourceToPbArticlesSource(entity entities.ArticlesSource) (*pb.ArticlesSource) {
	return &pb.ArticlesSource{
		Title: entity.Title,
		Description: entity.Description,
		Link: entity.Link,
		FeedLink: entity.FeedLink,
		Image: entity.Image,
	}
}

func CastArrayEntityArticleToPbType(entityArticles []*entities.Article) ([]*pb.Article) {
	pbArticles := make([]*pb.Article, 0)

	articlesByte, err := json.Marshal(entityArticles)
	if err != nil {
		log.Errorf("error occrus when marshal entity articles: %s", err)
	}

	err = json.Unmarshal(articlesByte, &pbArticles)
	if err != nil {
		log.Errorf("error occrus when unmarshal entity articles to pb article: %s", err)
	}
	return pbArticles
}
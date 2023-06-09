package helpers

import (
	log "github.com/sirupsen/logrus"
	"encoding/json"
	"server/entities"
	pb "server/proto"
)

func CastTestResult(result *pb.TestResult) (*entities.ArticlesSource, []*entities.Article) {
	articlesSource := CastPbArticleSourceToEntityArticlesSource(result.GetArticlesSource())
	articles := CastToPbArticlesToEntityArticles(result.GetArticles())
	return articlesSource, articles
}

func CastPbArticleSourceToEntityArticlesSource(pbArticlesSource *pb.ArticlesSource) (*entities.ArticlesSource) {
	return &entities.ArticlesSource{
		Title: pbArticlesSource.Title,
		Description: pbArticlesSource.Description,
		Link: pbArticlesSource.Link,
		FeedLink: pbArticlesSource.FeedLink,
		Image: pbArticlesSource.Image,
	}
}

func CastToPbArticlesToEntityArticles(pbArticles []*pb.Article) ([]*entities.Article) {
	articles := make([]*entities.Article, 0)

	articlesByte, err := json.Marshal(pbArticles)
	if err != nil {
		log.Errorf("error occrus when marshal entity articles: %s", err)
	}

	err = json.Unmarshal(articlesByte, &articles)
	if err != nil {
		log.Errorf("error occrus when unmarshal entity articles to pb article: %s", err)
	}
	return articles
}
package helpers

import (
	"server/services"
	"encoding/json"
	"server/entities"
	pb "server/proto"

	log "github.com/sirupsen/logrus"
)

func CastTestResult(result *pb.TestResult) (*services.ArticlesSourceResponse, []*services.ArticleResponse) {
	articlesSource := CastPbArticleSourceToArticlesSourceResponse(result.GetArticlesSource())
	articles := CastToPbArticlesToArticlesResponse(result.GetArticles())
	return articlesSource, articles
}

func CastPbArticleSourceToArticlesSourceResponse(pbArticlesSource *pb.ArticlesSource) (*services.ArticlesSourceResponse) {
	return &services.ArticlesSourceResponse{
		Title: pbArticlesSource.Title,
		Description: pbArticlesSource.Description,
		Link: pbArticlesSource.Link,
		FeedLink: pbArticlesSource.FeedLink,
		Image: pbArticlesSource.Image,
	}
}

func CastToPbArticlesToArticlesResponse(pbArticles []*pb.Article) ([]*services.ArticleResponse) {
	articles := make([]*services.ArticleResponse, 0)

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
package handlers

import (
	"crawler/entities"
	pb "crawler/proto"
	log "github.com/sirupsen/logrus"
	"encoding/json"
)

func castPbCrawlerToEntityCrawler(pbCrawler *pb.Crawler) (*entities.Crawler) {
	return &entities.Crawler{
		SourceLink: pbCrawler.SourceLink,
		FeedLink: pbCrawler.FeedLink,
		CrawlType: pbCrawler.CrawlType,
		ArticleDiv: pbCrawler.ArticleDiv,
		ArticleTitle: pbCrawler.ArticleTitle,
		ArticleDescription: pbCrawler.ArticleDescription,
		ArticleLink: pbCrawler.ArticleLink,
		ArticlePublished: pbCrawler.ArticlePublished,
		ArticleAuthors: pbCrawler.ArticleAuthors,
		Schedule: pbCrawler.Schedule,
	}
}

func newTestResult(articlesSource *entities.ArticlesSource, articles []*entities.Article) (*pb.TestResult){
	return &pb.TestResult{
		Articles: castArrayEntityArticleToPbType(articles),
		ArticlesSource: castEntityArticleSourceToPbArticlesSource(articlesSource),
	}
}

func castEntityArticleSourceToPbArticlesSource(entity *entities.ArticlesSource) (*pb.ArticlesSource) {
	return &pb.ArticlesSource{
		Title: entity.Title,
		Description: entity.Description,
		Link: entity.Link,
		FeedLink: entity.FeedLink,
		Image: entity.Image,
	}
}

func castArrayEntityArticleToPbType(entityArticles []*entities.Article) ([]*pb.Article) {
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
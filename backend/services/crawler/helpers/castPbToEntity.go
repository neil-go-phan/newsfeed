package helpers

import (
	"crawler/entities"
	pb "crawler/proto"
)

func CastPbCrawlerToEntityCrawler(pbCrawler *pb.Crawler) (entities.Crawler) {
	return entities.Crawler{
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
		ArticlesSourceID: uint(pbCrawler.ArticlesSourceId),
	}
}

func NewTestResult(articlesSource entities.ArticlesSource, articles []entities.Article) (*pb.TestResult){
	return &pb.TestResult{
		Articles: CastArrayEntityArticleToPbType(articles),
		ArticlesSource: CastEntityArticleSourceToPbArticlesSource(articlesSource),
	}
}

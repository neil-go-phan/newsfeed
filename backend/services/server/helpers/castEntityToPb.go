package helpers

import (
	"server/entities"
	pb "server/proto"
)

func CastEntityCrawlerToPbCrawler(entityCrawler entities.Crawler ) (*pb.Crawler) {
	return &pb.Crawler{
		SourceLink: entityCrawler.SourceLink,
		FeedLink: entityCrawler.FeedLink,
		CrawlType: entityCrawler.CrawlType,
		ArticleDiv: entityCrawler.ArticleDiv,
		ArticleTitle: entityCrawler.ArticleTitle,
		ArticleDescription: entityCrawler.ArticleDescription,
		ArticleLink: entityCrawler.ArticleLink,
		ArticleAuthors: entityCrawler.ArticleAuthors,
		Schedule: entityCrawler.Schedule,
		ArticlesSourceId: int32(entityCrawler.ArticlesSourceID),
	}
}
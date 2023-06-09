package crawlerservices

import (
	"context"
	"net/url"
	"server/entities"
	"server/helpers"
	pb "server/proto"
	"time"
)

func validateCrawler(crawler *entities.Crawler) (error) {
	_, err := url.ParseRequestURI(crawler.SourceLink)
	if err != nil {
		return  err
	}
	return nil
}

func getTestCrawlerResult(grpcClient pb.CrawlerServiceClient, crawler *entities.Crawler) (*pb.TestResult, error){
	ctx, cancle := context.WithTimeout(context.Background(), 2 * time.Minute)
	defer cancle()

	in := helpers.CastEntityCrawlerToPbCrawler(crawler)

	result, err := grpcClient.TestCrawler(ctx, in)
	if err != nil {
		return nil, err
	}
	return result, nil
}


package crawlerservices

import (
	"context"
	"fmt"
	"net/url"
	"server/entities"
	"server/helpers"
	pb "server/proto"
	"time"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

func validateCrawler(crawler *entities.Crawler) (error) {
	_, err := url.ParseRequestURI(crawler.SourceLink)
	if err != nil {
		return  err
	}
	return nil
}

func getTestCrawlerResult(grpcClient pb.CrawlerServiceClient, crawler *entities.Crawler) (*pb.TestResult, error){
	ctx, cancle := context.WithTimeout(context.Background(), 1 * time.Minute)
	defer cancle()

	in := helpers.CastEntityCrawlerToPbCrawler(crawler)

	result, err := grpcClient.TestCrawler(ctx, in)
	if err != nil {
		log.Error(err)
		status, _ := status.FromError(err)
		return nil, fmt.Errorf(status.Message()) 
	}
	return result, nil
}


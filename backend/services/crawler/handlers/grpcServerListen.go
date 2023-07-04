package handlers

import (
	"context"
	"net"

	log "github.com/sirupsen/logrus"

	"crawler/helpers"
	pb "crawler/proto"
	"crawler/services"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedCrawlerServiceServer
	crawlerService services.CrawlerServices
}

func NewGRPCServer(crawlerService services.CrawlerServices) *GRPCServer {
	return &GRPCServer{
		crawlerService: crawlerService,
	}
}

func (gRPC *GRPCServer) Listen(port string) {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	pb.RegisterCrawlerServiceServer(s, gRPC)
	log.Println("crawler gRPC server start listening")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (gRPC *GRPCServer) TestRSSCrawler(ctx context.Context, pbCrawler *pb.Crawler) (*pb.TestResult, error) {
	log.Println("Start test rss crawler")
	entityCrawler := helpers.CastPbCrawlerToEntityCrawler(pbCrawler)
	articlesSource, article, err := gRPC.crawlerService.TestRSSCrawler(entityCrawler)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	testResult := helpers.NewTestResult(articlesSource, article)

	log.Println("Test complete")
	return testResult, nil
}

func (gRPC *GRPCServer) TestCustomCrawler(ctx context.Context, pbCrawler *pb.Crawler) (*pb.TestResult, error) {
	log.Println("Start test custom crawler")
	entityCrawler := helpers.CastPbCrawlerToEntityCrawler(pbCrawler)
	articlesSource, article, err := gRPC.crawlerService.TestCustomCrawler(entityCrawler)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	testResult := helpers.NewTestResult(articlesSource, article)

	log.Println("Test complete")
	return testResult, nil
}

func (gRPC *GRPCServer) Crawl(ctx context.Context, pbCrawler *pb.Crawler) (*pb.NewArticlesCount, error) {
	log.Println("Start crawl",pbCrawler.SourceLink)
	entityCrawler := helpers.CastPbCrawlerToEntityCrawler(pbCrawler)

	newArticlesCount, err := gRPC.crawlerService.Crawl(entityCrawler)
	count := &pb.NewArticlesCount{
		Count: newArticlesCount,
	}

	if err != nil {
		log.Error(err)
		return count, err
	}

	return count, nil
}
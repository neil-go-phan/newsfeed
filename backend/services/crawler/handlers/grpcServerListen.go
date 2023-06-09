package handlers

import (
	"context"
	"net"

	log "github.com/sirupsen/logrus"

	pb "crawler/proto"
	"crawler/services"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedCrawlerServiceServer
	crawlerService services.CrawlerServices
}

func NewGRPCServer(crawlerService services.CrawlerServices) *GRPCServer{
	return &GRPCServer{
		crawlerService: crawlerService,
	}
}

func (gRPC *GRPCServer)Listen(port string) {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	pb.RegisterCrawlerServiceServer(s, &GRPCServer{})
	log.Println("crawler gRPC server start listening")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (gRPC *GRPCServer)TestCrawler(ctx context.Context, pbCrawler *pb.Crawler) (*pb.TestResult, error) {
	log.Println("Start test crawler")
	entityCrawler := castPbCrawlerToEntityCrawler(pbCrawler)
	
	articlesSource, article, err := gRPC.crawlerService.TestCrawler(entityCrawler)
	if err != nil {
		return nil, err
	}

	testResult := newTestResult(articlesSource, article)

	log.Println("Test complete")
	return testResult, nil
}


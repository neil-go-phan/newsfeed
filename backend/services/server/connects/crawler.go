package connects

import (
	"server/helpers"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectToCrawler(env helpers.EnvConfig) *grpc.ClientConn {
	conn, err := grpc.Dial(env.CrawlerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	return conn
}
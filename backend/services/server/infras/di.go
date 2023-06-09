package infras

import (
	pb "server/proto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func SetupRoute(db *gorm.DB, r *gin.Engine, grpcClient pb.CrawlerServiceClient) {
	userRoutes := InitizeUser(db)
	crawlerRoutes := InitizeCrawler(db, grpcClient)
	userRoutes.Setup(r)
	crawlerRoutes.Setup(r)
}
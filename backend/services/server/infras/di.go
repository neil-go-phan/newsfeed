package infras

import (
	"server/db/seed"
	pb "server/proto"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func SetupRoute(db *gorm.DB, r *gin.Engine, grpcClient pb.CrawlerServiceClient, jobIDMap map[string]cron.EntryID) {
	cronjob := cron.New()

	userRoutes := InitizeUser(db)
	crawlerRoutes := InitizeCrawler(db, grpcClient, cronjob, jobIDMap)
	topicRoutes := InitizeTopic(db)
	categoryRoutes := InitizeCategory(db)
	articlesSourceRoutes := InitizeArticlesSources(db)
	articlesRoutes := InitizeArticles(db)
	followRoutes := InitizeFollow(db)
	readRoutes := InitizeRead(db)
	readLaterRoutes := InitizeReadLater(db)
	roleRoutes := InitizeRole(db)
	permissionRoutes := InitizePermission(db)

	seed.Seed(db, grpcClient, jobIDMap)

	userRoutes.Setup(r)
	crawlerRoutes.Setup(r)
	topicRoutes.Setup(r)
	categoryRoutes.Setup(r)
	articlesSourceRoutes.Setup(r)
	articlesRoutes.Setup(r)
	followRoutes.Setup(r)
	readRoutes.Setup(r)
	readLaterRoutes.Setup(r)
	roleRoutes.Setup(r)
	permissionRoutes.Setup(r)

	// cronjob Setup
	go func() {
		crawlerRoutes.CreateCrawlerCronjobFromDB()
		cronjob.Run()
	}()
	
}

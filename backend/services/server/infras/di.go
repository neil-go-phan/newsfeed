package infras

import (
	"server/db/seed"
	pb "server/proto"

	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

func SetupRoute(db *gorm.DB, r *gin.Engine, grpcClient pb.CrawlerServiceClient, jobIDMap map[string]cron.EntryID) {
	cronjob := cron.New()
	fcmClient := createFCMClient()

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
	notificationRoutes := InitizeFcmNotification(db, fcmClient, cronjob)

	seed.Seed(db, grpcClient, jobIDMap, cronjob)

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
	notificationRoutes.Setup(r)

	// cronjob Setup
	go func() {
		// crawlerRoutes.CreateCrawlerCronjobFromDB()
		notificationRoutes.CreatePushNotificationCronjob()
		cronjob.Run()
	}()

}

func createFCMClient() *messaging.Client {
	opt := option.WithCredentialsFile("newfeed_firebase.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("new firebase app: %s", err)
	}
	fcmClient, err := app.Messaging(context.TODO())
	if err != nil {
		log.Fatalf("messaging: %s", err)
	}
	return fcmClient
}

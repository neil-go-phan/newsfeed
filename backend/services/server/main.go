package main

import (
	"fmt"
	"io"
	"os"
	"server/connects"
	"server/db/migrations"
	"server/helpers"
	"server/infras"
	"server/middlewares"
	pb "server/proto"
	"server/repository"
	"time"

	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"

	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"gorm.io/gorm"
)

func init() {
	// setup log
	log.SetLevel(log.InfoLevel)
	format := &easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg%\n",
	}
	log.SetFormatter(format)

	hook, err := logrus_sentry.NewSentryHook("https://0ed788b229564c98996c08db89759152@o4505503137857536.ingest.sentry.io/4505503139823616", []log.Level{
		log.PanicLevel,
		log.FatalLevel,
	})
	if err == nil {
		log.AddHook(hook)
	}
}

var (
	DB  *gorm.DB
	Gin *gin.Engine
	jobIDMap = make(map[string]cron.EntryID)
)

func main() {
	env, err := helpers.LoadEnv(".")
	if err != nil {
		log.Fatalln("cannot load env")
	}

	// write log file
	logFile, err := os.OpenFile("serverlog.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + "serverlog.log")
		panic(err)
	}
	defer logFile.Close()
	log.SetOutput(io.MultiWriter(logFile, os.Stdout))

	// sentry
	configSentry()
	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage("Connect to server success")

	// connect postgres and migration
	db := repository.ConnectDB(env.DBSource)
	migrations.RunDBMigration(env.MigrationURL, env.DBSource)

	// connect crawler
	conn := connects.ConnectToCrawler(env)
	grpcClient := pb.NewCrawlerServiceClient(conn)

	// app routes
	log.Infoln("Setup routes")
	r := gin.Default()
	r.Use(middlewares.Cors())
	
	infras.SetupRoute(db, r, grpcClient, jobIDMap)

	err = r.Run(env.Port)
	if err != nil {
		log.Fatalln("error occurred when run server")
	}
}

func configSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://dc16289c8fd744249297e6150d4bc9dc@o4505040225501184.ingest.sentry.io/4505277257285632",
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}

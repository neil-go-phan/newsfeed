package main

import (
	"crawler/helpers"
	"crawler/repository"
	"fmt"
	"io"
	"os"
	"crawler/infras"
	"time"

	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
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
		log.ErrorLevel,
	})
	if err == nil {
		log.AddHook(hook)
	}
}

func main() {
	env, err := helpers.LoadEnv(".")
	if err != nil {
		log.Fatalln("cannot load env: ", err)
	}

	logFile, err := os.OpenFile("crawlerlog.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + "crawlerlog.log")
		panic(err)
	}
	defer logFile.Close()

	log.SetOutput(io.MultiWriter(logFile, os.Stdout))

	configSentry()
	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage("Connect to crawler success")


	db := repository.ConnectDB(env.DBSource)

	infras.Listen(db, env.GRPCPort)
}

func configSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://4cad04fffc3348dc8d14d1f592f1d014@o4505040225501184.ingest.sentry.io/4505066672947200",
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}

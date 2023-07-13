package notificationservices

import (
	"context"
	"server/entities"
	"server/repository"
	"server/services"

	"firebase.google.com/go/messaging"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

type FcmNotificationService struct {
	repo           repository.FcmNotificationRepository
	articleService services.ArticleServices
	fcmClient      *messaging.Client
	cron           *cron.Cron
}

const CHUNKS_SIZE = 500
const MESSAGE_TITLE = "Most read article"

func NewFcmNotificationService(repo repository.FcmNotificationRepository, articleService services.ArticleServices, fcmClient *messaging.Client, cron *cron.Cron) *FcmNotificationService {
	service := &FcmNotificationService{
		repo:           repo,
		articleService: articleService,
		fcmClient:      fcmClient,
		cron:           cron,
	}
	return service
}

func (s *FcmNotificationService) List() ([]entities.FcmNotification, error) {
	return s.repo.List()
}

func (s *FcmNotificationService) Create(fcmNotification entities.FcmNotification) error {
	return s.repo.Create(fcmNotification)
}

func (s *FcmNotificationService) SendMessages(message *messaging.MulticastMessage) {
	batchResp, err := s.fcmClient.SendMulticast(context.TODO(), message)
	log.Printf("batch response: %+v, err: %s \n", batchResp, err)
}

func (s *FcmNotificationService) SendArticleToAll() {
	log.Println("start send notification...")
	list, err := s.repo.List()
	if err != nil {
		log.Error(err)
	}
	article, err := s.articleService.GetMostReadInHour()
	if err != nil {
		log.Error(err)
	}
	chunks := chunkSlice(list)
	for _, chunk := range chunks {
		tokens := getTokes(chunk)
		message := &messaging.MulticastMessage{
			Notification: &messaging.Notification{
				Title: MESSAGE_TITLE,
				Body: article.Title,
			},
			Tokens: tokens,
		}
		s.SendMessages(message)
	}
	log.Println("finish send notification...")
}

func (s *FcmNotificationService) CronjobPushNotification() {
	_, err := s.cron.AddFunc("@every 0h1m", func() { s.SendArticleToAll() })
	if err != nil {
		log.Error("error occurred while seting up cronjob: ", err)
	}

	log.Printf("create cronjob send article to all user hourly")
}

func chunkSlice(slice []entities.FcmNotification) [][]entities.FcmNotification {
	var chunks [][]entities.FcmNotification
	for i := 0; i < len(slice); i += CHUNKS_SIZE {
		end := i + CHUNKS_SIZE
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func getTokes(slice []entities.FcmNotification) []string {
	var listToken []string
	for _, item := range slice {
		if item.FirebaseToken != "" {
			listToken = append(listToken, item.FirebaseToken)
		}
		
	}
	return listToken
}

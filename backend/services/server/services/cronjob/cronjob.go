package cronjobservices

import (
	"fmt"
	"server/entities"
	pb "server/proto"
	"server/repository"
	"server/services"
	"time"

	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

// each crawler will have only one cronjob in runtime
// jobIDMap: map[crawler_url]cron.EntryID use to track cronjob with crawler url
type CronjobService struct {
	repo       repository.CronjobRepository
	cron       *cron.Cron
	grpcClient pb.CrawlerServiceClient
	jobIDMap   map[string]cron.EntryID
}

func NewCronjobService(repo repository.CronjobRepository, cronjob *cron.Cron, grpcClient pb.CrawlerServiceClient, jobIDMap map[string]cron.EntryID) *CronjobService {
	CronjobService := &CronjobService{
		repo:           repo,
		cron:           cronjob,
		grpcClient:     grpcClient,
		jobIDMap:       jobIDMap,
	}
	return CronjobService
}

func (s *CronjobService) CreateCrawlerCronjob(crawler entities.Crawler) {
	cronjobName := fmt.Sprintf("%s, %s", crawler.SourceLink, crawler.Schedule)
	entryID, err := s.cron.AddFunc(crawler.Schedule, func() { s.cronjobCrawlerFunction(crawler, cronjobName) })
	if err != nil {
		log.Error("error occurred while seting up cronjob: ", err)
	}

	log.Printf("create cronjob %s", cronjobName)
	mapKey := newMapKey(crawler.SourceLink, crawler.Schedule)
	s.jobIDMap[mapKey] = entryID
}

func (s *CronjobService) cronjobCrawlerFunction(crawler entities.Crawler, cronjobName string) {
	log.Println("create a new record on cronjob table")
	cronjobDB, err := s.createCronjobCrawlerDB(crawler, cronjobName)
	if err != nil {
		log.Error(err)
	}
	log.Println("start crawl")
	
	newArticleCount, err := crawl(crawler, s.grpcClient)
	if err != nil {
		log.Error(err)
	}

	log.Println("update end_at collumn in cronjob table")
	cronjobDB.NewArticlesCount = newArticleCount
	err = s.updateResult(cronjobDB)
	if err != nil {
		log.Error(err)
	}
}

func (s *CronjobService) createCronjobCrawlerDB(crawler entities.Crawler, cronjobName string) (*entities.Cronjob, error) {
	cronjobDB := &entities.Cronjob{
		StartAt:     time.Now(),
		CrawlerID:   crawler.ID,
		Crawler:     crawler,
		Name:        cronjobName,
		
	}

	cronjobDB, err := s.repo.Create(cronjobDB)
	if err != nil {
		return nil, err
	}
	return cronjobDB, nil
}

func (s *CronjobService) updateResult(cronjob *entities.Cronjob) error {
	cronjob.EndAt = time.Now()
	err := s.repo.UpdateResult(cronjob)
	if err != nil {
		return err
	}
	return nil
}

func (s *CronjobService) GetCronjobRuntime() []services.CronjobResponse {
	cronjobResponses := make([]services.CronjobResponse, 0)
	for key := range s.jobIDMap {
		log.Println(key)
		cronjobResponses = append(cronjobResponses, newCronjobResponse(key))
	}
	return cronjobResponses
}

func (s *CronjobService) CronjobOnHour(timeString string) (*[60]services.ChartHour, error) {
	chartsData := createArrayCronjobInHour()

	cronjobs, err := getCronjobsRunOnHour(timeString, s.repo)
	if err != nil {
		return &chartsData, err
	}

	chartsData = fillHourChartData(*cronjobs, chartsData)

	return &chartsData, nil
}

func (s *CronjobService) CronjobOnDay(timeString string) (*[24]services.ChartDay, error) {
	chartsData := createArrayCronjobOnDay()

	cronjobs, err := getCronjobsRunInDay(timeString, s.repo)
	if err != nil {
		return &chartsData, err
	}

	chartsData = fillDayChartData(*cronjobs, chartsData)


	return &chartsData, nil
}

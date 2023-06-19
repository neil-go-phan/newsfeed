package cronjobservices

import (
	"context"
	"fmt"
	"server/entities"
	pb "server/proto"
	"server/repository"
	"server/services"
	"strings"
	"time"
	// log "github.com/sirupsen/logrus"
)

const DEFAULT_YEAR = 1

func CastEntityCrawlerToPbCrawler(entityCrawler entities.Crawler) *pb.Crawler {
	return &pb.Crawler{
		SourceLink:         entityCrawler.SourceLink,
		FeedLink:           entityCrawler.FeedLink,
		CrawlType:          entityCrawler.CrawlType,
		ArticleDiv:         entityCrawler.ArticleDiv,
		ArticleTitle:       entityCrawler.ArticleTitle,
		ArticleDescription: entityCrawler.ArticleDescription,
		ArticleLink:        entityCrawler.ArticleLink,
		ArticlePublished:   entityCrawler.ArticlePublished,
		ArticleAuthors:     entityCrawler.ArticleAuthors,
		Schedule:           entityCrawler.Schedule,
		ArticlesSourceId: int32(entityCrawler.ArticlesSourceID),
	}
}

func newCronjobResponse(mapKey string) services.CronjobResponse {
	if !strings.Contains(mapKey, "$") {
		return services.CronjobResponse{
			Name: mapKey,
			Url:  "none",
		}
	}
	part := strings.Split(mapKey, "$")
	cronjobResponse := services.CronjobResponse{
		Name: part[0],
		Url:  part[1],
	}
	return cronjobResponse
}

func newMapKey(url string, schedule string) string {
	cronjobName := fmt.Sprintf("%s, %s", url, schedule)
	return fmt.Sprintf("%s$%s$%s", cronjobName, url, schedule)
}

func newCronjobInChart(cronjob entities.Cronjob) services.CronjobInChart {
	startAt := castTimeToString(cronjob.StartedAt)

	var endAt string
	if cronjob.EndedAt.Year() != DEFAULT_YEAR {
		endAt = castTimeToString(cronjob.EndedAt)
	} else {
		endAt = "runing"
	}

	return services.CronjobInChart{
		Name:    cronjob.Name,
		StartAt: startAt,
		EndAt:   endAt,
	}
}

func castTimeToString(time time.Time) string {
	hour := time.Hour()
	min := time.Minute()
	str := fmt.Sprintf("%v:%v", addZeroWhenLowwerThanTen(hour), addZeroWhenLowwerThanTen(min))
	return str
}

func addZeroWhenLowwerThanTen(time int) string {
	if time < 10 {
		return fmt.Sprintf("0%v", time)
	}
	return fmt.Sprintf("%v", time)
}

func createArrayCronjobInHour() [60]services.ChartHour {
	charts := [60]services.ChartHour{}
	for index := range charts {
		charts[index].Minute = index
		charts[index].AmountOfJob = 0
	}
	return charts
}

func getCronjobsRunOnHour(timeString string, repo repository.CronjobRepository) (entities *[]entities.Cronjob, err error) {
	hour, endOfHour, err := readHourFromFrontend(timeString)
	if err != nil {
		return entities, err
	}

	entities, err = repo.Get(hour, endOfHour)
	if err != nil {
		return entities, err
	}
	return entities, nil
}

func readHourFromFrontend(timeString string) (hour time.Time, endOfHour time.Time, err error) {
	hour, err = time.Parse("2006-01-02 15", timeString)
	if err != nil {
		return hour, endOfHour, err
	}

	endOfHour = hour.Add(time.Duration(59) * time.Minute)
	return hour, endOfHour, nil
}

func fillHourChartData(cronjobs []entities.Cronjob, charts [60]services.ChartHour) [60]services.ChartHour{
	minuteNow := time.Now().Minute()
	for _, entityChart := range cronjobs {

		minuteStart := entityChart.StartedAt.Minute()
		if entityChart.EndedAt.Year() != DEFAULT_YEAR {
			minuteEnd := entityChart.EndedAt.Minute()
			for i := minuteStart; i <= minuteEnd; i++ {
				charts[i].AmountOfJob += 1
				charts[i].Cronjobs = append(charts[i].Cronjobs, newCronjobInChart(entityChart))
			}
		} else {
			minuteEnd := minuteNow
			for i := minuteStart; i <= minuteEnd; i++ {
				charts[i].AmountOfJob += 1
				charts[i].Cronjobs = append(charts[i].Cronjobs, newCronjobInChart(entityChart))
			}
		}
	}

	return charts
}

func createArrayCronjobOnDay() [24]services.ChartDay {
	charts := [24]services.ChartDay{}
	for index := range charts {
		charts[index].Hour = index
		charts[index].AmountOfJob = 0
	}
	return charts
}

func getCronjobsRunInDay(timeString string, repo repository.CronjobRepository) (entities *[]entities.Cronjob, err error) {
	day, endOfDay, err := readDayFromFrontend(timeString)
	if err != nil {
		return entities, err
	}

	entities, err = repo.Get(day, endOfDay)
	if err != nil {
		return entities, err
	}
	return entities, nil
}

func readDayFromFrontend(timeString string) (day time.Time, endOfDay time.Time, err error) {
	day, err = time.Parse("2006-01-02", timeString)
	if err != nil {
		return day, endOfDay, err
	}

	endOfDay = day.Add(time.Duration(23)*time.Hour + time.Duration(59)*time.Minute)
	return day, endOfDay, nil
}

func fillDayChartData(cronjobs []entities.Cronjob, charts [24]services.ChartDay) [24]services.ChartDay{
	for _, entityChart := range cronjobs {
		hour := entityChart.StartedAt.Hour()
		charts[hour].AmountOfJob += 1
		if charts[hour].Cronjobs == nil {
			charts[hour].Cronjobs = map[string]int{}
		}
		charts[hour].Cronjobs[entityChart.Name] += 1
	}

	return charts
}

func crawl(crawler entities.Crawler, grpcClient pb.CrawlerServiceClient) (int32, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 1 * time.Minute)
	defer cancle()

	pbCrawler := CastEntityCrawlerToPbCrawler(crawler)

	pbNewArticleCount, err := grpcClient.Crawl(ctx, pbCrawler)
	if err != nil {
		return 0, err 
	}
	
	return pbNewArticleCount.Count, nil
}
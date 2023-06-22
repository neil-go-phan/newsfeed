package services

import (
	"server/entities"
	"time"
)

type LoginUserInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserChangePasswordInput struct {
	Username             string `json:"username" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}

type RegisterUserInput struct {
	Username             string `json:"username" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
	Email                string `json:"email" validate:"required"`
}

type GoogleOauthToken struct {
	Access_token string
	Id_token     string
}

type GoogleUserResult struct {
	Id             string
	Email          string
	Verified_email bool
	Name           string
	Given_name     string
	Family_name    string
	Picture        string
	Locale         string
}

type ArticlesSourceResponseCrawl struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	FeedLink    string `json:"feed_link"`
	Image       string `json:"image"` // link or a base64 image
}

type ArticleResponse struct {
	ID          uint    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Published   time.Time `json:"published"`
	Authors     string    `json:"authors"`
}

type ArticlesSourceFromFrontend struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	FeedLink    string `json:"feed_link"`
	Image       string `json:"image"` // base64 image
	TopicID     uint   `json:"topic_id"`
}

type CrawlerFromFrontend struct {
	SourceLink         string `json:"source_link" validate:"required,url"`
	FeedLink           string `json:"feed_link"`
	CrawlType          string `json:"crawl_type" validate:"required"`
	ArticleDiv         string `json:"article_div"`
	ArticleTitle       string `json:"article_title"`
	ArticleDescription string `json:"article_description"`
	ArticleLink        string `json:"article_link"`
	ArticleAuthors     string `json:"article_authors"`
	// Schedule           string `json:"schedule"`
}

type CreateCrawlerPayload struct {
	ArticlesSource ArticlesSourceFromFrontend `json:"articles_source"`
	Crawler        CrawlerFromFrontend        `json:"crawler"`
}

type CronjobResponse struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Schedule string `json:"schedulte"`
}

type CronjobInChart struct {
	Name    string `json:"name"`
	StartAt string `json:"start_at"` // ex: "16:01"
	EndAt   string `json:"end_at"`   // ex: "16:02"
}

type ChartHour struct {
	Minute      int              `json:"minute"`
	AmountOfJob int              `json:"amount_of_jobs"`
	Cronjobs    []CronjobInChart `json:"cronjobs"`
}

type ChartDay struct {
	Hour        int            `json:"hour"`
	AmountOfJob int            `json:"amount_of_jobs"`
	Cronjobs    map[string]int // map[cronjob_name]runnng_times
}

type UpdateNameCategoryPayload struct {
	Category entities.Category `json:"category"`
	NewName  string            `json:"new_name"`
}

type CategoryResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Illustration string `json:"illustration"`
}

type TopicResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	CategoryID uint   `json:"category_id"`
}

type ArticlesSourceResponseRender struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Image       string `json:"image"` // base64 image
	TopicID     uint   `json:"topic_id"`
}

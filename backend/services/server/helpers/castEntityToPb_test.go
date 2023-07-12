package helpers

import (
	"server/entities"
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestCastEntityCrawlerToPbCrawler(t *testing.T) {
	entityCrawler := entities.Crawler{
		SourceLink:          "https://example.com",
		FeedLink:            "https://example.com/feed",
		CrawlType:           "feed",
		ArticleDiv:          "div.article",
		ArticleTitle:        "h1.title",
		ArticleDescription:  "p.description",
		ArticleLink:         "a.link",
		ArticleAuthors:      "John Doe, Jane Smith",
		Schedule:            "0 0 * * *",
		ArticlesSourceID:    123,
	}
	pbCrawler := CastEntityCrawlerToPbCrawler(entityCrawler)

	assert.Equal(t, entityCrawler.SourceLink, pbCrawler.SourceLink)
	assert.Equal(t, entityCrawler.FeedLink, pbCrawler.FeedLink)
	assert.Equal(t, entityCrawler.CrawlType, pbCrawler.CrawlType)
	assert.Equal(t, entityCrawler.ArticleDiv, pbCrawler.ArticleDiv)
	assert.Equal(t, entityCrawler.ArticleTitle, pbCrawler.ArticleTitle)
	assert.Equal(t, entityCrawler.ArticleDescription, pbCrawler.ArticleDescription)
	assert.Equal(t, entityCrawler.ArticleLink, pbCrawler.ArticleLink)
	assert.Equal(t, entityCrawler.ArticleAuthors, pbCrawler.ArticleAuthors)
	assert.Equal(t, entityCrawler.Schedule, pbCrawler.Schedule)
	assert.Equal(t, int32(entityCrawler.ArticlesSourceID), pbCrawler.ArticlesSourceId)

}
package helpers

import (
	"server/entities"
	pb "server/proto"
	"server/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCastTestResult(t *testing.T) {
	pbArticlesSource := &pb.ArticlesSource{
		Title:       "Test Title",
		Description: "Test Description",
		Link:        "https://example.com",
		FeedLink:    "https://example.com/feed",
		Image:       "https://example.com/image.jpg",
	}

	pbArticles := []*pb.Article{
		{
			Title:       "Article 1",
			Description: "Content 1",
			Link:        "link1.com",
			Authors:     "author1",
		},
		{
			Title:       "Article 2",
			Description: "Content 2",
			Link:        "link2.com",
			Authors:     "author1",
		},
	}

	testResult := &pb.TestResult{
		ArticlesSource: pbArticlesSource,
		Articles:       pbArticles,
	}

	articlesSource, articles := CastTestResult(testResult)

	expectedArticlesSource := &services.ArticlesSourceResponseCrawl{
		Title:       "Test Title",
		Description: "Test Description",
		Link:        "https://example.com",
		FeedLink:    "https://example.com/feed",
		Image:       "https://example.com/image.jpg",
	}
	assert.Equal(t, expectedArticlesSource, articlesSource)

	expectedArticles := []*services.ArticleResponse{
		{
			Title:       "Article 1",
			Description: "Content 1",
			Link:        "link1.com",
			Authors:     "author1",
		},
		{
			Title:       "Article 2",
			Description: "Content 2",
			Link:        "link2.com",
			Authors:     "author1",
		},
	}
	assert.Equal(t, expectedArticles, articles)
}

func TestCastPbArticleSourceToEntityArticlesSource(t *testing.T) {
	pbArticlesSource := &pb.ArticlesSource{
		Title:       "Test Title",
		Description: "Test Description",
		Link:        "https://example.com",
		FeedLink:    "https://example.com/feed",
		Image:       "https://example.com/image.jpg",
	}

	articlesSource := CastPbArticleSourceToEntityArticlesSource(pbArticlesSource)

	expectedArticlesSource := &entities.ArticlesSource{
		Title:       "Test Title",
		Description: "Test Description",
		Link:        "https://example.com",
		FeedLink:    "https://example.com/feed",
		Image:       "https://example.com/image.jpg",
	}
	assert.Equal(t, expectedArticlesSource, articlesSource)
}

func TestCastToPbArticlesToEntityArticles(t *testing.T) {
	pbArticles := []*pb.Article{
		{
			Title:       "Article 1",
			Description: "Content 1",
			Link:        "link1.com",
			Authors:     "author1",
		},
		{
			Title:       "Article 2",
			Description: "Content 2",
			Link:        "link2.com",
			Authors:     "author1",
		},
	}

	articles := CastToPbArticlesToEntityArticles(pbArticles)

	expectedArticles := []*entities.Article{
		{
			Title:       "Article 1",
			Description: "Content 1",
			Link:        "link1.com",
			Authors:     "author1",
		},
		{
			Title:       "Article 2",
			Description: "Content 2",
			Link:        "link2.com",
			Authors:     "author1",
		},
	}
	assert.Equal(t, expectedArticles, articles)
}

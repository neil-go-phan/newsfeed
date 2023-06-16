package articleservices

import (
	"crawler/entities"
	"crawler/repository"
)

type ArticleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) *ArticleService {
	articleService := &ArticleService{
		repo: repo,
	}
	return articleService
}

func (s *ArticleService) CreateIfNotExist(article entities.Article) error {
	return s.repo.CreateIfNotExist(article)
}

func (s *ArticleService) StoreArticles(articles []entities.Article, articlesSourceID uint) (count int32) {
	for _, article := range articles {
		article.ArticlesSourceID = articlesSourceID
		err := s.repo.CreateIfNotExist(article)
		if err == nil {
			count++
		}
	}
	return count
}

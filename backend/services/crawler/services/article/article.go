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

func (s *ArticleService) CreateIfNotExist(article *entities.Article) (error) {
	return s.repo.CreateIfNotExist(article) 
}

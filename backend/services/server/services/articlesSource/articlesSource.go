package articlessourceservices

import (
	"server/entities"
	"server/repository"
)

type ArticlesSourceService struct {
	repo repository.ArticlesSourcesRepository
}

func NewArticlesSourceService(repo repository.ArticlesSourcesRepository) *ArticlesSourceService {
	articlesSourceService := &ArticlesSourceService{
		repo: repo,
	}
	return articlesSourceService
}

func (s *ArticlesSourceService) Create(articlesSource *entities.ArticlesSource) (error) {
	return s.repo.Create(articlesSource) 
}

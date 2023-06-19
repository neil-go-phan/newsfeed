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

func (s *ArticlesSourceService) CreateIfNotExist(articlesSource entities.ArticlesSource) (entities.ArticlesSource, error) {
	return s.repo.CreateIfNotExist(articlesSource)
}

func (s *ArticlesSourceService) UpdateTopicOneSource(articlesSource entities.ArticlesSource, newTopicId uint) error {
	return s.repo.UpdateTopicOneSource(articlesSource, newTopicId)
}

func (s *ArticlesSourceService) UpdateTopicAllSource(oldTopicId uint, newTopicId uint) error {
	return s.repo.UpdateTopicAllSource(oldTopicId, newTopicId)
}

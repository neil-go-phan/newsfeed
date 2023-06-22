package articlessourceservices

import (
	"server/entities"
	"server/repository"
	"server/services"
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

func (s *ArticlesSourceService) GetByTopicID(topicID uint) ([]services.ArticlesSourceResponseRender, error) {
	articlesSourcesResponse := make([]services.ArticlesSourceResponseRender, 0)
	articlesSources, err := s.repo.GetWithTopic(topicID)
	if err != nil {
		return articlesSourcesResponse, err
	}
	for _, articlesSource := range articlesSources {
		articlesSourcesResponse = append(articlesSourcesResponse, castEntityArticlesSourceToReponse(articlesSource))

	}
	return articlesSourcesResponse, nil
}

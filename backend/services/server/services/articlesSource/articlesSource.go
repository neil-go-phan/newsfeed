package articlessourceservices

import (
	"server/entities"
	"server/repository"
	"server/services"

	"gorm.io/gorm"
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

func (s *ArticlesSourceService) GetByTopicIDPaginate(topicID uint, page int, pageSize int) ([]services.ArticlesSourceResponseRender, int64, error) {
	articlesSourcesResponse := make([]services.ArticlesSourceResponseRender, 0)
	articlesSources, found, err := s.repo.GetWithTopicPaginate(topicID, page, pageSize)
	if err != nil {
		return articlesSourcesResponse, found, err
	}
	for _, articlesSource := range articlesSources {
		articlesSourcesResponse = append(articlesSourcesResponse, castEntityArticlesSourceToReponse(articlesSource))

	}
	return articlesSourcesResponse, found, nil
}

func (s *ArticlesSourceService) SearchByTitleAndDescriptionPaginate(keyword string, page int, pageSize int) ([]services.ArticlesSourceResponseRender, int64, error) {
	articlesSourcesResponse := make([]services.ArticlesSourceResponseRender, 0)
	articlesSources, found, err := s.repo.SearchByTitleAndDescriptionPaginate(keyword, page, pageSize)
	if err != nil {
		return articlesSourcesResponse, found, err
	}
	for _, articlesSource := range articlesSources {
		articlesSourcesResponse = append(articlesSourcesResponse, castEntityArticlesSourceToReponse(articlesSource))

	}
	return articlesSourcesResponse, found, nil
}

func (s *ArticlesSourceService) UserFollow(articlesSourceID uint) error {
	articlesSource := entities.ArticlesSource{
		Model: gorm.Model{ID: articlesSourceID},
	}
	return s.repo.IncreaseFollowByOne(articlesSource)
}

func (s *ArticlesSourceService) UserUnfollow(articlesSourceID uint) error {
	articlesSource := entities.ArticlesSource{
		Model: gorm.Model{ID: articlesSourceID},
	}
	return s.repo.DecreaseFollowByOne(articlesSource)
}

func (s *ArticlesSourceService) GetMostActiveSources() ([]services.ArticlesSourceRecommended, error) {
	articlesSourcesResponse := make([]services.ArticlesSourceRecommended, 0)
	articlesSources, err := s.repo.GetMostActiveSources()
	if err != nil {
		return articlesSourcesResponse, err
	}
	for _, articlesSource := range articlesSources {
		articlesSourcesResponse = append(articlesSourcesResponse, castArticlesSourceRecommended(articlesSource))

	}
	return articlesSourcesResponse, nil
}

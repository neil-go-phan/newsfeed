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

func (s *ArticlesSourceService) Search(keyword string, page int, pageSize int) ([]services.ArticlesSourceResponseRender, int64, error) {
	articlesSourcesResponse := make([]services.ArticlesSourceResponseRender, 0)
	articlesSources, found, err := s.repo.Search(keyword, page, pageSize)
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

func (s *ArticlesSourceService) ListAll() ([]services.ArticlesSourceResponseRender, error) {
	articlesSourcesResponse := make([]services.ArticlesSourceResponseRender, 0)
	articlesSources, err := s.repo.ListAll()
	if err != nil {
		return articlesSourcesResponse, err
	}
	for _, articlesSource := range articlesSources {
		articlesSourcesResponse = append(articlesSourcesResponse, castEntityArticlesSourceToReponse(articlesSource))

	}
	return articlesSourcesResponse, nil
}

func (s *ArticlesSourceService) GetWithID(id uint) (services.ArticlesSourceResponseRender, error) {
	articlesSourcesResponse := services.ArticlesSourceResponseRender{}
	articlesSource, err := s.repo.GetWithID(id)
	if err != nil {
		return articlesSourcesResponse, err
	}
	articlesSourcesResponse = castEntityArticlesSourceToReponse(articlesSource)
	return articlesSourcesResponse, nil
}

func (s *ArticlesSourceService) Count() (int, error) {
	return s.repo.Count()
}

func (s *ArticlesSourceService) ListAllPaging(page int, pageSize int) ([]services.ArticlesSourceResponseRender, error) {
	articlesSourcesResponse := make([]services.ArticlesSourceResponseRender, 0)
	articlesSources, err := s.repo.ListAllPaging(page, pageSize)
	if err != nil {
		return articlesSourcesResponse, err
	}
	for _, articlesSource := range articlesSources {
		articlesSourcesResponse = append(articlesSourcesResponse, castEntityArticlesSourceToReponse(articlesSource))

	}
	return articlesSourcesResponse, nil
}

func (s *ArticlesSourceService) SearchWithFilter(keyword string, page int, pageSize int, topicID uint) ([]services.ArticlesSourceResponseRender, int64, error) {
	articlesSourcesResponse := make([]services.ArticlesSourceResponseRender, 0)

	articlesSources, found, err := s.searchOptions(keyword, page, pageSize, topicID)
	if err != nil {
		return articlesSourcesResponse, found, err
	}
	for _, source := range articlesSources {
		articlesSourcesResponse = append(articlesSourcesResponse, castEntityArticlesSourceToReponse(source))

	}
	return articlesSourcesResponse, found, nil
}

func (s *ArticlesSourceService) searchOptions(keyword string, page int, pageSize int, topicID uint) ([]entities.ArticlesSource, int64, error) {
	// without key word
	if keyword == "" {
		return s.repo.GetWithTopicPaginate(topicID, page, pageSize)
	}
	if topicID == 0 {
		return s.repo.Search(keyword, page, pageSize)
	}
	return s.repo.SearchWithFilter(keyword, page, pageSize, topicID)
}

func (s *ArticlesSourceService) Delete(sourceID uint) error {
	source := entities.ArticlesSource{
		Model: gorm.Model{
			ID: sourceID,
		},
	}
	return s.repo.Delete(source)
}

func (s *ArticlesSourceService) Update(articlesSource entities.ArticlesSource) error {
	return s.repo.Update(articlesSource)
}

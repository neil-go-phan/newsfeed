package topicservices

import (
	"server/entities"
	"server/repository"
	"server/services"
	"strings"
)

const OTHERS_TOPIC_ID = 1
const OTHERS_TOPIC_NAME = "Others"

type TopicService struct {
	repo                    repository.TopicRepository
	articlesSourcesServices services.ArticlesSourceServices
}

func NewTopicService(repo repository.TopicRepository, articlesSourcesServices services.ArticlesSourceServices) *TopicService {
	topicService := &TopicService{
		repo:                    repo,
		articlesSourcesServices: articlesSourcesServices,
	}
	return topicService
}

func (s *TopicService) CreateIfNotExist(topic entities.Topic) error {
	topic.Name = strings.TrimSpace(topic.Name)
	err := validateTopic(topic)
	if err != nil {
		return err
	}
	return s.repo.CreateIfNotExist(topic)
}

func (s *TopicService) List() ([]services.TopicResponse, error) {
	topicsResponse := make([]services.TopicResponse, 0)
	topics, err := s.repo.List()
	if err != nil {
		return topicsResponse, err
	}
	for _, topic := range topics {
		topicsResponse = append(topicsResponse, castEntityTopicToTopicResponse(topic))
	}
	return topicsResponse, nil
}

func (s *TopicService) UpdateWhenDeteleCategory(oldCategoryID uint, newCategoryID uint) error {
	return s.repo.UpdateWhenDeteleCategory(oldCategoryID, newCategoryID)
}

func (s *TopicService) Update(topic entities.Topic) error {
	topic.Name = strings.TrimSpace(topic.Name)
	err := validateTopic(topic)
	if err != nil {
		return err
	}
	return s.repo.Update(topic)
}

func (s *TopicService) Delete(topic entities.Topic) error {
	topic.Name = strings.TrimSpace(topic.Name)
	err := validateTopic(topic)
	if err != nil {
		return err
	}

	err = s.articlesSourcesServices.UpdateTopicAllSource(topic.ID, OTHERS_TOPIC_ID)
	if err != nil {
		return err
	}
	return s.repo.Delete(topic)
}

func (s *TopicService) GetPagination(page int, pageSize int) ([]services.TopicResponse, error) {
	topicsResponse := make([]services.TopicResponse, 0)
	topics, err := s.repo.GetPagination(page, pageSize)
	if err != nil {
		return topicsResponse, err
	}
	for _, topic := range topics {
		topicsResponse = append(topicsResponse, castEntityTopicToTopicResponse(topic))
	}
	return topicsResponse, nil
}

func (s *TopicService) Count() (int, error) {
	return s.repo.Count()
}

func (s *TopicService) GetByCategory(categoryID uint) ([]services.TopicResponse, error) {
	topicsResponse := make([]services.TopicResponse, 0)
	topics, err := s.repo.GetByCategory(categoryID)
	if err != nil {
		return topicsResponse, err
	}
	for _, topic := range topics {
		topicsResponse = append(topicsResponse, castEntityTopicToTopicResponse(topic))
	}
	return topicsResponse, nil
}

func (s *TopicService) SearchByName(keyword string) ([]services.TopicResponse, error) {
	topicsResponse := make([]services.TopicResponse, 0)
	topics, err := s.repo.SearchByName(keyword)
	if err != nil {
		return topicsResponse, err
	}
	for _, topic := range topics {
		topicsResponse = append(topicsResponse, castEntityTopicToTopicResponse(topic))
	}
	return topicsResponse, nil
}

func (s *TopicService) SearchTopicAndArticlesSourcePaginate(keyword string, page int, pageSize int) ([]services.TopicResponse, []services.ArticlesSourceResponseRender, int64, error) {
	topics, err := s.SearchByName(keyword)
	if err != nil {
		return []services.TopicResponse{}, []services.ArticlesSourceResponseRender{}, 0, err
	}

	articlesSources, articleSourcesFound, err := s.articlesSourcesServices.Search(keyword, page, pageSize)
	if err != nil {
		return []services.TopicResponse{}, []services.ArticlesSourceResponseRender{}, articleSourcesFound, err
	}

	return topics, articlesSources, articleSourcesFound, nil
}

package followservices

import (
	"server/entities"
	"server/repository"
	"server/services"
)

type FollowService struct {
	repo                  repository.FollowRepository
	articleSourceServices services.ArticlesSourceServices
}

func NewFollowService(repo repository.FollowRepository, articleSourceServices services.ArticlesSourceServices) *FollowService {
	followService := &FollowService{
		repo:                  repo,
		articleSourceServices: articleSourceServices,
	}
	return followService
}

func (s *FollowService) Follow(username string, articlesSourceID uint) error {
	follow := entities.Follow{
		Username:         username,
		ArticlesSourceID: articlesSourceID,
	}
	err := s.repo.CreateIfNotExist(follow)
	if err != nil {
		return err
	}

	return s.articleSourceServices.UserFollow(articlesSourceID)
}

func (s *FollowService) Unfollow(username string, articlesSourceID uint) error {
	follow := entities.Follow{
		Username:         username,
		ArticlesSourceID: articlesSourceID,
	}
	err := s.repo.Delete(follow)
	if err != nil {
		return err
	}
	return s.articleSourceServices.UserUnfollow(articlesSourceID)
}

func (s *FollowService) GetUserFollowedSources(username string) ([]services.ArticlesSourceUserFollow, error) {
	follows, err := s.repo.GetByUsername(username)
	if err != nil {
		return []services.ArticlesSourceUserFollow{}, err
	}
	articlesSources := getArticlesSourcesFromArrayFollow(follows)

	return articlesSources, nil
}

func (s *FollowService) GetNewestSourceUpdatedID(username string) ([]uint, error) {
	follows, err := s.repo.GetNewestFeedsUpdated(username)
	if err != nil {
		return []uint{}, err
	}
	ids := getArticlesSourceID(follows)

	return ids, nil
}
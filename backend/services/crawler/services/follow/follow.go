package followservices

import (
	"crawler/repository"
)

type FollowService struct {
	repo                  repository.FollowRepository
}

func NewFollowService(repo repository.FollowRepository) *FollowService {
	followService := &FollowService{
		repo:                  repo,
	}
	return followService
}

func (s *FollowService) UpdateNewUnreadArticles(sourceID uint, countNewArticleFound int32) error {
	if (countNewArticleFound > 0) {
		return s.repo.UpdateUnreadBySourceID(sourceID, countNewArticleFound)
	}
	return nil 
}
package readservices

import (
	"server/entities"
	"server/repository"
	"server/services"
)

type ReadService struct {
	repo repository.ReadRepository
	followServices services.FollowServices
}

func NewReadService(repo repository.ReadRepository, followServices services.FollowServices) *ReadService {
	ReadService := &ReadService{
		repo: repo,
		followServices: followServices,
	}
	return ReadService
}

func (s *ReadService) MarkArticleAsRead(username string, articleID uint, articlesSourceID uint) error {
	read := entities.Read{
		Username: username,
		ArticleID: articleID,
		ArticlesSourceID: articlesSourceID,
	}

	err := s.repo.Create(read)
	if err != nil {
		return err
	}

	return nil
}

func (s *ReadService) MarkArticleAsUnRead(username string, articleID uint, articlesSourceID uint) error {
	read := entities.Read{
		Username: username,
		ArticleID: articleID,
		ArticlesSourceID: articlesSourceID,
	}
	err := s.repo.Delete(read)
	if err != nil {
		return err
	}

	return nil
}

func (s *ReadService) CountAllUnreadArticles(username string, articleID uint, articlesSourceID uint) error {
	read := entities.Read{
		Username: username,
		ArticleID: articleID,
		ArticlesSourceID: articlesSourceID,
	}
	return s.repo.Delete(read)
}
package readservices

import (
	"server/entities"
	"server/repository"
)

type ReadService struct {
	repo repository.ReadRepository
}

func NewReadService(repo repository.ReadRepository) *ReadService {
	ReadService := &ReadService{
		repo: repo,
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

func (s *ReadService) MarkAllAsReadBySourceID(username string, articlesSourceID uint) error {
	return s.repo.MarkAllAsReadBySourceID(username, articlesSourceID)
}

func (s *ReadService) MarkAllAsReadByUserFollowedSource(username string) error {
	return s.repo.MarkAllAsReadByUserFollowedSource(username)
}
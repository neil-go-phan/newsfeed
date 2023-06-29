package readlaterservices

import (
	"server/entities"
	"server/repository"
)

type ReadLaterService struct {
	repo repository.ReadLaterRepository	
}

func NewReadLaterService(repo repository.ReadLaterRepository) *ReadLaterService {
	readLaterService := &ReadLaterService{
		repo: repo,
	}
	return readLaterService
}

func (s *ReadLaterService) AddToReadLaterList(username string, articleID uint) error {
	readLater := entities.ReadLater{
		Username: username,
		ArticleID: articleID,
	}

	err := s.repo.Create(readLater)
	if err != nil {
		return err
	}

	return nil
}

func (s *ReadLaterService) RemoveFromReadLaterList(username string, articleID uint) error {
	readLater := entities.ReadLater{
		Username: username,
		ArticleID: articleID,
	}

	err := s.repo.Delete(readLater)
	if err != nil {
		return err
	}

	return nil
}
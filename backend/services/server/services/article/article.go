package articleservices

import (
	"server/entities"
	"server/repository"
	"server/services"
)

type ArticleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) *ArticleService {
	articleService := &ArticleService{
		repo: repo,
	}
	return articleService
}

func (s *ArticleService) CreateIfNotExist(article *entities.Article) error {
	return s.repo.CreateIfNotExist(article)
}

// func (s *ArticleService) GetArticlesPaginationByArticlesSourceID(articlesSourceID uint, page int, pageSize int) ([]services.ArticleResponse, error) {
// 	// articlesResponse := make([]services.ArticleResponse, 0)
// 	// articles, err := s.repo.GetArticlesPaginationByArticlesSourceID(articlesSourceID, page, pageSize)
// 	// if err != nil {
// 	// 	return articlesResponse, err
// 	// }
// 	// for _, articles := range articles {
// 	// 	articlesResponse = append(articlesResponse, castEntityArticleToReponse(articles))

// 	// }
// 	// return articlesResponse, nil
// }

func (s *ArticleService) GetArticlesPaginationByArticlesSourceID(username string, articlesSourceID uint, page int, pageSize int) ([]services.ArticleForReadResponse, error) {
	articlesResponse := make([]services.ArticleForReadResponse, 0)
	articles, err := s.repo.GetArticlesPaginationByArticlesSourceIDWithReadStatus(username, articlesSourceID, page, pageSize)
	if err != nil {
		return articlesResponse, err
	}
	for _, article := range articles {
		articlesResponse = append(articlesResponse, castArticleFromRepoToArticleReadReponse(article))

	}
	return articlesResponse, nil
}

func (s *ArticleService) GetArticlesPaginationByUserFollowedSource(username string, page int, pageSize int) ([]services.ArticleForReadResponse, error) {
	articlesResponse := make([]services.ArticleForReadResponse, 0)
	articles, err := s.repo.GetArticlesPaginationByUserFollowedSource(username, page, pageSize)
	if err != nil {
		return articlesResponse, err
	}
	for _, article := range articles {
		articlesResponse = append(articlesResponse, castArticleFromRepoToArticleReadReponse(article))

	}
	return articlesResponse, nil
}	

func (s *ArticleService) GetUnreadArticlesPaginationByArticlesSourceID(username string, articlesSourceID uint, page int, pageSize int) ([]services.ArticleForReadResponse, error) {
	articlesResponse := make([]services.ArticleForReadResponse, 0)
	articles, err := s.repo.GetUnreadArticlesPaginationByArticlesSourceID(username, articlesSourceID, page, pageSize)
	if err != nil {
		return articlesResponse, err
	}
	for _, article := range articles {
		articlesResponse = append(articlesResponse, castArticleFromRepoToArticleReadReponse(article))

	}
	return articlesResponse, nil
}

func (s *ArticleService) GetUnreadArticlesByUserFollowedSource(username string, page int, pageSize int) ([]services.ArticleForReadResponse, error) {
	articlesResponse := make([]services.ArticleForReadResponse, 0)
	articles, err := s.repo.GetUnreadArticlesByUserFollowedSource(username, page, pageSize)
	if err != nil {
		return articlesResponse, err
	}
	for _, article := range articles {
		articlesResponse = append(articlesResponse, castArticleFromRepoToArticleReadReponse(article))

	}
	return articlesResponse, nil
}	

func (s *ArticleService) CountArticleCreateAWeekAgoByArticlesSourceID(articlesSourceID uint) (int64, error) {
	return s.repo.CountArticleCreateAWeekAgoByArticlesSourceID(articlesSourceID)
}

func (s *ArticleService) SearchArticlesAcrossUserFollowedSources(username string,keyword string, page int, pageSize int) ([]services.ArticleResponse,int64, error) {
	articlesResponse := make([]services.ArticleResponse, 0)
	articles,found, err := s.repo.SearchArticlesAcrossUserFollowedSources(username,keyword, page, pageSize)
	if err != nil {
		return articlesResponse,found, err
	}
	for _, articles := range articles {
		articlesResponse = append(articlesResponse, castEntityArticleToReponse(articles))

	}
	return articlesResponse,found, nil
}

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

func (s *ArticleService) GetPaginationByArticlesSourceID(articlesSourceID uint, page int, pageSize int) ([]services.ArticleResponse, error) {
	articlesResponse := make([]services.ArticleResponse, 0)
	articles, err := s.repo.GetPaginationByArticlesSourceID(articlesSourceID, page, pageSize)
	if err != nil {
		return articlesResponse, err
	}
	for _, articles := range articles {
		articlesResponse = append(articlesResponse, castEntityArticleToReponse(articles))

	}
	return articlesResponse, nil
}

func (s *ArticleService) GetPaginationByUserFollowedSource(username string, page int, pageSize int) ([]services.ArticleResponse, error) {
	articlesResponse := make([]services.ArticleResponse, 0)
	articles, err := s.repo.GetPaginationByUserFollowedSource(username, page, pageSize)
	if err != nil {
		return articlesResponse, err
	}
	for _, articles := range articles {
		articlesResponse = append(articlesResponse, castEntityArticleToReponse(articles))

	}
	return articlesResponse, nil
}	

func (s *ArticleService) SearchArticlesAcrossSources(keyword string, page int, pageSize int) ([]services.ArticleResponse,int64, error) {
	articlesResponse := make([]services.ArticleResponse, 0)
	articles,found, err := s.repo.SearchArticlesAcrossSources(keyword, page, pageSize)
	if err != nil {
		return articlesResponse,found, err
	}
	for _, articles := range articles {
		articlesResponse = append(articlesResponse, castEntityArticleToReponse(articles))

	}
	return articlesResponse,found, nil
}

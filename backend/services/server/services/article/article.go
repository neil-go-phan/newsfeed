package articleservices

import (
	"server/entities"
	"server/repository"
	"server/services"
	"time"

	"gorm.io/gorm"
)

type ArticleService struct {
	repo         repository.ArticleRepository
	roleServices services.RoleServices
}

func NewArticleService(repo repository.ArticleRepository, roleServices services.RoleServices) *ArticleService {
	articleService := &ArticleService{
		repo:         repo,
		roleServices: roleServices,
	}
	return articleService
}

func (s *ArticleService) CreateIfNotExist(article *entities.Article) error {
	return s.repo.CreateIfNotExist(article)
}

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

func (s *ArticleService) GetUnreadArticlesPaginationByArticlesSourceID( username string, articlesSourceID uint, page int, pageSize int) ([]services.ArticleForReadResponse, error) {
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

func (s *ArticleService) GetUnreadArticlesByUserFollowedSource( username string, page int, pageSize int) ([]services.ArticleForReadResponse, error) {
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

func (s *ArticleService) GetReadLaterListPaginationByArticlesSourceID(username string, articlesSourceID uint, page int, pageSize int) ([]services.ArticleForReadResponse, error) {
	articlesResponse := make([]services.ArticleForReadResponse, 0)

	articles, err := s.repo.GetReadLaterListPaginationByArticlesSourceID(username, articlesSourceID, page, pageSize)
	if err != nil {
		return articlesResponse, err
	}
	for _, article := range articles {
		articlesResponse = append(articlesResponse, castArticleFromRepoToArticleReadReponse(article))

	}
	return articlesResponse, nil
}

func (s *ArticleService) GetReadLaterListPaginationByUserFollowedSource(username string, page int, pageSize int) ([]services.ArticleForReadResponse, error) {
	articlesResponse := make([]services.ArticleForReadResponse, 0)

	articles, err := s.repo.GetReadLaterListPaginationByUserFollowedSource(username, page, pageSize)
	if err != nil {
		return articlesResponse, err
	}
	for _, article := range articles {
		articlesResponse = append(articlesResponse, castArticleFromRepoToArticleReadReponse(article))

	}
	return articlesResponse, nil
}

func (s *ArticleService) GetRecentlyReadArticle(username string, page int, pageSize int) ([]services.ArticleForReadResponse, error) {
	articlesResponse := make([]services.ArticleForReadResponse, 0)
	articles, err := s.repo.GetRecentlyReadArticle(username, page, pageSize)
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

func (s *ArticleService) SearchArticlesAcrossUserFollowedSources(username string, keyword string, page int, pageSize int) ([]services.ArticleResponse, int64, error) {
	articlesResponse := make([]services.ArticleResponse, 0)
	articles, found, err := s.repo.SearchArticlesAcrossUserFollowedSources(username, keyword, page, pageSize)
	if err != nil {
		return articlesResponse, found, err
	}
	for _, articles := range articles {
		articlesResponse = append(articlesResponse, castEntityArticleToReponse(articles))

	}
	return articlesResponse, found, nil
}

func (s *ArticleService) GetTredingArticle(username string) ([]services.TredingArticleResponse, error) {
	articlesResponse := make([]services.TredingArticleResponse, 0)
	articles, err := s.repo.GetTredingArticle(username)
	if err != nil {
		return articlesResponse, err
	}
	for _, article := range articles {
		articlesResponse = append(articlesResponse, castTrendingArticle(article))

	}
	return articlesResponse, nil
}

func (s *ArticleService) ListAll(page int, pageSize int) ([]services.ArticleResponse, error) {
	articlesResponse := make([]services.ArticleResponse, 0)
	articles, err := s.repo.ListAll(page, pageSize)
	if err != nil {
		return articlesResponse, err
	}
	for _, articles := range articles {
		articlesResponse = append(articlesResponse, castEntityArticleToReponse(articles))

	}
	return articlesResponse, nil
}

func (s *ArticleService) Count() (int, error) {
	return s.repo.Count()
}

func (s *ArticleService) Delete(articleID uint) error {
	article := entities.Article{
		Model: gorm.Model{
			ID: articleID,
		},
	}
	return s.repo.Delete(article)
}

func (s *ArticleService) AdminSearchArticlesWithFilter(keyword string, page int, pageSize int, articlesSourceID uint) ([]services.ArticleResponse, int64, error) {
	articlesResponse := make([]services.ArticleResponse, 0)

	articles, found, err := s.searchOptions(keyword, page, pageSize, articlesSourceID)
	if err != nil {
		return articlesResponse, found, err
	}
	for _, articles := range articles {
		articlesResponse = append(articlesResponse, castEntityArticleToReponse(articles))

	}
	return articlesResponse, found, nil
}

func (s *ArticleService) searchOptions(keyword string, page int, pageSize int, articlesSourceID uint) ([]entities.Article, int64, error) {
	// without key word
	if keyword == "" {
		return s.repo.GetArticlesPaginationByArticlesSourceID(articlesSourceID, page, pageSize)
	}
	if articlesSourceID == 0 {
		return s.repo.AdminSearchArticles(keyword, page, pageSize)
	}
	return s.repo.AdminSearchArticlesWithFilter(keyword, page, pageSize, articlesSourceID)
}

func (s *ArticleService) GetMostReadInDay() (entities.Article, error) {
	now := time.Now()
	previousDay := now.Add(time.Duration(-24) * time.Hour)
	return s.repo.GetMostRead(previousDay, now)
}
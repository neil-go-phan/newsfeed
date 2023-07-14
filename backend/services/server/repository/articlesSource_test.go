package repository

import (
	"fmt"
	"regexp"
	"server/entities"
	"strings"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

const TEST_TOPIC_ID = 1
const TEST_ARTICLES_SOURCE_ID = 1

func (s *Suite) TestArticlesSources_GetWithTopicPaginate() {
	repo := NewArticlesSourcesRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "articles_sources" WHERE topic_id = $1 AND "articles_sources"."deleted_at" IS NULL LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_TOPIC_ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "nothing"))

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "articles_sources" WHERE topic_id = $1 AND "articles_sources"."deleted_at" IS NULL LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_TOPIC_ID).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	_, _, err := repo.GetWithTopicPaginate(TEST_TOPIC_ID, PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticlesSources_CreateIfNotExist() {
	repo := NewArticlesSourcesRepo(s.DB)

	articleSources := entities.ArticlesSource{
		Title:       "title test1",
		Description: "description test1",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "articles_sources" WHERE "articles_sources"."deleted_at" IS NULL ORDER BY "articles_sources"."id" LIMIT 1`)).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description"}).
			AddRow("title test1", "description test1"))

	_, err := repo.CreateIfNotExist(articleSources)
	s.Assert().EqualError(err, "article source already exist")
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticlesSources_Search() {
	repo := NewArticlesSourcesRepo(s.DB)

	keyword := "keyword"
	searchKeyword := fmt.Sprint("%" + strings.ToLower(keyword) + "%")

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "articles_sources" WHERE (LOWER(title) LIKE $1 or LOWER(description) LIKE $2 or LOWER(link) LIKE $3) AND "articles_sources"."deleted_at" IS NULL LIMIT 10 OFFSET 10`)).
		WithArgs(searchKeyword, searchKeyword, searchKeyword).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description"}).
			AddRow("title test1", "description test1"))

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "articles_sources" WHERE (LOWER(title) LIKE $1 or LOWER(description) LIKE $2 or LOWER(link) LIKE $3) AND "articles_sources"."deleted_at" IS NULL LIMIT 10 OFFSET 10`)).
		WithArgs(searchKeyword, searchKeyword, searchKeyword).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	_, _, err := repo.Search(keyword, PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticlesSources_GetMostActiveSources() {
	repo := NewArticlesSourcesRepo(s.DB)

	today := time.Now()
	aWeekAgo := today.AddDate(0, 0, -7)
	todayString := today.Format("2006-01-02") + " 23:59:59"
	aWeekAgoString := aWeekAgo.Format("2006-01-02") + " 00:00:00"

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT DISTINCT "id","title","description","link","follower","image",articles_previous_week,"feed_link" FROM "articles_sources" JOIN (SELECT "articles_source_id",count(id) as articles_previous_week FROM "articles" WHERE (created_at BETWEEN $1 AND $2) AND "articles"."deleted_at" IS NULL GROUP BY "articles_source_id" LIMIT 5) q on q.articles_source_id = articles_sources.id WHERE "articles_sources"."deleted_at" IS NULL ORDER BY articles_previous_week desc`)).
		WithArgs(aWeekAgoString, todayString).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description"}).
			AddRow("title test1", "description test1"))

	_, err := repo.GetMostActiveSources()
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticlesSources_ListAll() {
	repo := NewArticlesSourcesRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "articles_sources" WHERE "articles_sources"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description"}).
			AddRow("title test1", "description test1"))

	_, err := repo.ListAll()
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticlesSources_GetWithID() {
	repo := NewArticlesSourcesRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "articles_sources" WHERE id = $1 AND "articles_sources"."deleted_at" IS NULL ORDER BY "articles_sources"."id" LIMIT 1`)).
		WithArgs(TEST_ARTICLES_SOURCE_ID).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description"}).
			AddRow("title test1", "description test1"))

	_, err := repo.GetWithID(TEST_ARTICLES_SOURCE_ID)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticlesSources_Count() {
	repo := NewArticlesSourcesRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "articles_sources"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	_, err := repo.Count()
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticlesSources_ListAllPaging() {
	repo := NewArticlesSourcesRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "articles_sources" WHERE "articles_sources"."deleted_at" IS NULL ORDER BY created_at desc LIMIT 10 OFFSET 10`)).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description"}).
			AddRow("title test1", "description test1"))

	_, err := repo.ListAllPaging(PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticlesSources_SearchWithFilter() {
	repo := NewArticlesSourcesRepo(s.DB)

	keyword := "keyword"
	searchKeyword := fmt.Sprint("%" + strings.ToLower(keyword) + "%")

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "articles_sources" WHERE (topic_id = $1 AND (LOWER(title) LIKE $2 or LOWER(description) LIKE $3 or LOWER(link) LIKE $4)) AND "articles_sources"."deleted_at" IS NULL LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_TOPIC_ID, searchKeyword, searchKeyword, searchKeyword).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description"}).
			AddRow("title test1", "description test1"))

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "articles_sources" WHERE (topic_id = $1 AND (LOWER(title) LIKE $2 or LOWER(description) LIKE $3 or LOWER(link) LIKE $4)) AND "articles_sources"."deleted_at" IS NULL LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_TOPIC_ID, searchKeyword, searchKeyword, searchKeyword).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	_, _, err := repo.SearchWithFilter(keyword, PAGE, PAGE_SIZE, TEST_TOPIC_ID)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

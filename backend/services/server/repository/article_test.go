package repository

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
}

const TEST_USERNAME = "username"
const PAGE = 2
const PAGE_SIZE = 10
const TEST_SOURCE_ID = 1

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	require.NoError(s.T(), err)

	conn, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	s.DB = conn
	require.NoError(s.T(), err)
}

func TestInit(t *testing.T) {
	s := new(Suite)
	suite.Run(t, s)
}

func (s *Suite) TestArticles_GetArticlesPaginationByArticlesSourceID() {
	repo := NewArticleRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "articles" WHERE articles_source_id = $1 AND "articles"."deleted_at" IS NULL ORDER BY created_at desc LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_SOURCE_ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "nothing"))

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "articles" WHERE articles_source_id = $1 AND "articles"."deleted_at" IS NULL LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_SOURCE_ID).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	repo.GetArticlesPaginationByArticlesSourceID(TEST_SOURCE_ID, PAGE, PAGE_SIZE)

	err := s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_GetArticlesPaginationByArticlesSourceIDWithReadStatus() {
	repo := NewArticleRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT DISTINCT "id","title","description","link","published","authors",articles.articles_source_id,articles.created_at,CASE WHEN r.username IS NULL THEN false ELSE true END AS is_read,CASE WHEN rl.username IS NULL THEN false ELSE true END AS is_read_later FROM "articles" LEFT JOIN (SELECT * FROM "reads" WHERE username = $1 AND articles_source_id = $2) r on articles.id = r.article_id LEFT JOIN (SELECT * FROM "read_laters" WHERE username = $3) rl on articles.id = rl.article_id WHERE articles.articles_source_id = $4 AND "articles"."deleted_at" IS NULL ORDER BY articles.created_at desc LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_USERNAME, TEST_SOURCE_ID, TEST_USERNAME, TEST_SOURCE_ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "nothing"))

	repo.GetArticlesPaginationByArticlesSourceIDWithReadStatus(TEST_USERNAME, TEST_SOURCE_ID, PAGE, PAGE_SIZE)

	err := s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_GetArticlesPaginationByUserFollowedSource() {
	repo := NewArticleRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT DISTINCT "id","title","description","link","published","authors",articles.articles_source_id,articles.created_at,CASE WHEN r.username IS NULL THEN false ELSE true END AS is_read,CASE WHEN rl.username IS NULL THEN false ELSE true END AS is_read_later FROM "articles" JOIN (SELECT "articles_source_id" FROM "follows" WHERE username = $1) f on f.articles_source_id = articles.articles_source_id LEFT JOIN reads r on articles.id = r.article_id LEFT JOIN (SELECT username, article_id FROM "read_laters" WHERE username = $2) rl on articles.id = rl.article_id WHERE "articles"."deleted_at" IS NULL ORDER BY articles.created_at desc LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_USERNAME, TEST_USERNAME).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "nothing"))

	repo.GetArticlesPaginationByUserFollowedSource(TEST_USERNAME, PAGE, PAGE_SIZE)

	err := s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_GetReadLaterListPaginationByArticlesSourceID() {
	repo := NewArticleRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT DISTINCT "id","title","description","link","published","authors",articles.articles_source_id,articles.created_at,CASE WHEN r.username IS NULL THEN false ELSE true END AS is_read,CASE WHEN rl.username IS NULL THEN false ELSE true END AS is_read_later FROM "articles" LEFT JOIN (SELECT username, article_id FROM "reads" WHERE username = $1 AND articles_source_id = $2) r on articles.id = r.article_id JOIN (SELECT username, article_id FROM "read_laters" WHERE username = $3) rl on articles.id = rl.article_id WHERE articles.articles_source_id = $4 AND "articles"."deleted_at" IS NULL ORDER BY articles.created_at desc LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_USERNAME, TEST_SOURCE_ID, TEST_USERNAME, TEST_SOURCE_ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "nothing"))

	repo.GetReadLaterListPaginationByArticlesSourceID(TEST_USERNAME, TEST_SOURCE_ID, PAGE, PAGE_SIZE)

	err := s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_GetReadLaterListPaginationByUserFollowedSource() {
	repo := NewArticleRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT DISTINCT "id","title","description","link","published","authors",articles.articles_source_id,articles.created_at,CASE WHEN r.username IS NULL THEN false ELSE true END AS is_read,CASE WHEN rl.username IS NULL THEN false ELSE true END AS is_read_later FROM "articles" JOIN (SELECT "articles_source_id" FROM "follows" WHERE username = $1) f on f.articles_source_id = articles.articles_source_id LEFT JOIN reads r on articles.id = r.article_id JOIN (SELECT username, article_id FROM "read_laters" WHERE username = $2) rl on articles.id = rl.article_id WHERE "articles"."deleted_at" IS NULL ORDER BY articles.created_at desc LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_USERNAME, TEST_USERNAME).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "nothing"))

	repo.GetReadLaterListPaginationByUserFollowedSource(TEST_USERNAME, PAGE, PAGE_SIZE)

	err := s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}
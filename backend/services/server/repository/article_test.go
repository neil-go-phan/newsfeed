package repository

import (
	"database/sql"
	"fmt"
	"regexp"
	"server/entities"
	"strings"
	"testing"
	"time"

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

	_, _, err := repo.GetArticlesPaginationByArticlesSourceID(TEST_SOURCE_ID, PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_GetArticlesPaginationByArticlesSourceIDWithReadStatus() {
	repo := NewArticleRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT DISTINCT "id","title","description","link","published","authors",articles.articles_source_id,articles.created_at,CASE WHEN r.username IS NULL THEN false ELSE true END AS is_read,CASE WHEN rl.username IS NULL THEN false ELSE true END AS is_read_later FROM "articles" LEFT JOIN (SELECT * FROM "reads" WHERE username = $1 AND articles_source_id = $2) r on articles.id = r.article_id LEFT JOIN (SELECT * FROM "read_laters" WHERE username = $3) rl on articles.id = rl.article_id WHERE articles.articles_source_id = $4 AND "articles"."deleted_at" IS NULL ORDER BY articles.created_at desc LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_USERNAME, TEST_SOURCE_ID, TEST_USERNAME, TEST_SOURCE_ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "nothing"))

	_, err := repo.GetArticlesPaginationByArticlesSourceIDWithReadStatus(TEST_USERNAME, TEST_SOURCE_ID, PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_GetArticlesPaginationByUserFollowedSource() {
	repo := NewArticleRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT DISTINCT "id","title","description","link","published","authors",articles.articles_source_id,articles.created_at,CASE WHEN r.username IS NULL THEN false ELSE true END AS is_read,CASE WHEN rl.username IS NULL THEN false ELSE true END AS is_read_later FROM "articles" JOIN (SELECT "articles_source_id" FROM "follows" WHERE username = $1) f on f.articles_source_id = articles.articles_source_id LEFT JOIN reads r on articles.id = r.article_id LEFT JOIN (SELECT username, article_id FROM "read_laters" WHERE username = $2) rl on articles.id = rl.article_id WHERE "articles"."deleted_at" IS NULL ORDER BY articles.created_at desc LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_USERNAME, TEST_USERNAME).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "nothing"))

	_, err := repo.GetArticlesPaginationByUserFollowedSource(TEST_USERNAME, PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_GetReadLaterListPaginationByArticlesSourceID() {
	repo := NewArticleRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT DISTINCT "id","title","description","link","published","authors",articles.articles_source_id,articles.created_at,CASE WHEN r.username IS NULL THEN false ELSE true END AS is_read,CASE WHEN rl.username IS NULL THEN false ELSE true END AS is_read_later FROM "articles" LEFT JOIN (SELECT username, article_id FROM "reads" WHERE username = $1 AND articles_source_id = $2) r on articles.id = r.article_id JOIN (SELECT username, article_id FROM "read_laters" WHERE username = $3) rl on articles.id = rl.article_id WHERE articles.articles_source_id = $4 AND "articles"."deleted_at" IS NULL ORDER BY articles.created_at desc LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_USERNAME, TEST_SOURCE_ID, TEST_USERNAME, TEST_SOURCE_ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "nothing"))

	_, err := repo.GetReadLaterListPaginationByArticlesSourceID(TEST_USERNAME, TEST_SOURCE_ID, PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_GetReadLaterListPaginationByUserFollowedSource() {
	repo := NewArticleRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT DISTINCT "id","title","description","link","published","authors",articles.articles_source_id,articles.created_at,CASE WHEN r.username IS NULL THEN false ELSE true END AS is_read,CASE WHEN rl.username IS NULL THEN false ELSE true END AS is_read_later FROM "articles" LEFT JOIN reads r on articles.id = r.article_id JOIN (SELECT username, article_id FROM "read_laters" WHERE username = $1) rl on articles.id = rl.article_id WHERE "articles"."deleted_at" IS NULL ORDER BY articles.created_at desc LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_USERNAME).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "nothing"))

	_, err := repo.GetReadLaterListPaginationByUserFollowedSource(TEST_USERNAME, PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_GetUnreadArticlesPaginationByArticlesSourceID() {
	repo := NewArticleRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT DISTINCT "id","title","description","link","published","authors",articles.articles_source_id,articles.created_at FROM "articles" LEFT OUTER JOIN (SELECT * FROM "reads" WHERE username = $1 AND articles_source_id = $2) q on articles.id = q.article_id WHERE (articles.articles_source_id = $3 AND q.article_id IS NULL) AND "articles"."deleted_at" IS NULL ORDER BY articles.created_at desc LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_USERNAME, TEST_SOURCE_ID, TEST_SOURCE_ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "nothing"))

	_, err := repo.GetUnreadArticlesPaginationByArticlesSourceID(TEST_USERNAME, TEST_SOURCE_ID, PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_GetUnreadArticlesByUserFollowedSource() {
	repo := NewArticleRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT DISTINCT "id","title","description","link","published","authors",articles.articles_source_id,articles.created_at FROM "articles" JOIN (SELECT "articles_source_id" FROM "follows" WHERE username = $1) q on q.articles_source_id = articles.articles_source_id LEFT OUTER JOIN reads r on articles.id = r.article_id WHERE r.article_id IS NULL AND "articles"."deleted_at" IS NULL ORDER BY articles.created_at desc LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_USERNAME).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "nothing"))

	_, err := repo.GetUnreadArticlesByUserFollowedSource(TEST_USERNAME, PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_GetRecentlyReadArticle() {
	repo := NewArticleRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT DISTINCT "id","title","description","link","published","authors",articles.articles_source_id,r.created_at,CASE WHEN r.article_id IS NULL THEN false ELSE true END AS is_read FROM "articles" JOIN (SELECT "articles_source_id" FROM "follows" WHERE username = $1) q on q.articles_source_id = articles.articles_source_id JOIN reads r on articles.id = r.article_id WHERE "articles"."deleted_at" IS NULL ORDER BY r.created_at desc LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_USERNAME).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).
			AddRow(1, "nothing"))

	_, err := repo.GetRecentlyReadArticle(TEST_USERNAME, PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_CountArticleCreateAWeekAgoByArticlesSourceID() {
	repo := NewArticleRepo(s.DB)
	today := time.Now()
	aWeekAgo := today.AddDate(0, 0, -7)
	todayString := today.Format("2006-01-02") + " 23:59:59"
	aWeekAgoString := aWeekAgo.Format("2006-01-02") + " 00:00:00"

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "articles" WHERE (articles_source_id = $1 AND created_at BETWEEN $2 AND $3) AND "articles"."deleted_at" IS NULL`)).
		WithArgs(TEST_SOURCE_ID, aWeekAgoString, todayString).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	_, err := repo.CountArticleCreateAWeekAgoByArticlesSourceID(TEST_SOURCE_ID)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_SearchArticlesAcrossUserFollowedSources() {
	repo := NewArticleRepo(s.DB)
	keyword := "Test"
	searchKeyword := fmt.Sprint("%" + strings.ToLower(keyword) + "%")

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT DISTINCT "title","description","link","published","authors",q.articles_source_id,"created_at" FROM "articles" JOIN (SELECT "articles_source_id" FROM "follows" WHERE username = $1) q on articles.articles_source_id = q.articles_source_id WHERE (LOWER(title) LIKE $2 or LOWER(description) LIKE $3) AND "articles"."deleted_at" IS NULL ORDER BY created_at desc LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_USERNAME, searchKeyword, searchKeyword).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description"}).
			AddRow("title test1", "description test1"))

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "articles" JOIN (SELECT "articles_source_id" FROM "follows" WHERE username = $1) q on articles.articles_source_id = q.articles_source_id WHERE (LOWER(title) LIKE $2 or LOWER(description) LIKE $3) AND "articles"."deleted_at" IS NULL LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_USERNAME, searchKeyword, searchKeyword).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	_, _, err := repo.SearchArticlesAcrossUserFollowedSources(TEST_USERNAME, keyword, PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_AdminSearchArticles() {
	repo := NewArticleRepo(s.DB)
	keyword := "Test"
	searchKeyword := fmt.Sprint("%" + strings.ToLower(keyword) + "%")

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT DISTINCT "title","description","link","published","authors","articles_source_id","created_at" FROM "articles" WHERE (LOWER(title) LIKE $1 or LOWER(description) LIKE $2) AND "articles"."deleted_at" IS NULL ORDER BY created_at desc LIMIT 10 OFFSET 10`)).
		WithArgs(searchKeyword, searchKeyword).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description"}).
			AddRow("title test1", "description test1"))

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "articles" WHERE (LOWER(title) LIKE $1 or LOWER(description) LIKE $2) AND "articles"."deleted_at" IS NULL LIMIT 10 OFFSET 10`)).
		WithArgs(searchKeyword, searchKeyword).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	_, _, err := repo.AdminSearchArticles(keyword, PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_AdminSearchArticlesWithFilter() {
	repo := NewArticleRepo(s.DB)
	keyword := "Test"
	searchKeyword := fmt.Sprint("%" + strings.ToLower(keyword) + "%")

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT DISTINCT "title","description","link","published","authors","articles_source_id","created_at" FROM "articles" WHERE (articles_source_id = $1 AND (LOWER(title) LIKE $2 or LOWER(description) LIKE $3)) AND "articles"."deleted_at" IS NULL ORDER BY created_at desc LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_SOURCE_ID, searchKeyword, searchKeyword).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description"}).
			AddRow("title test1", "description test1"))

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "articles" WHERE (articles_source_id = $1 AND (LOWER(title) LIKE $2 or LOWER(description) LIKE $3)) AND "articles"."deleted_at" IS NULL LIMIT 10 OFFSET 10`)).
		WithArgs(TEST_SOURCE_ID, searchKeyword, searchKeyword).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	_, _, err := repo.AdminSearchArticlesWithFilter(keyword, PAGE, PAGE_SIZE, TEST_SOURCE_ID)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_GetTredingArticle() {
	repo := NewArticleRepo(s.DB)
	today := time.Now()
	todayString := today.Format("2006-01-02")
	from := todayString + " 00:00:00"
	to := todayString + " 23:59:59"

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id","title","description","link","published","authors",articles.articles_source_id,articles.created_at,CASE WHEN rl.username IS NULL THEN false ELSE true END AS is_read_later FROM "articles" JOIN (SELECT "article_id",count(article_id) as read FROM "reads" WHERE created_at between $1 AND $2 GROUP BY "article_id" ORDER BY read desc LIMIT 10) r on articles.id = r.article_id LEFT JOIN (SELECT * FROM "read_laters" WHERE username = $3) rl on articles.id = rl.article_id WHERE "articles"."deleted_at" IS NULL ORDER BY r.read desc`)).
		WithArgs(from, to, TEST_USERNAME).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description"}).
			AddRow("title test1", "description test1"))

	_, err := repo.GetTredingArticle(TEST_USERNAME)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_ListAll() {
	repo := NewArticleRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "articles" WHERE "articles"."deleted_at" IS NULL ORDER BY created_at desc LIMIT 10 OFFSET 10`)).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description"}).
			AddRow("title test1", "description test1"))

	_, err := repo.ListAll(PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_Count() {
	repo := NewArticleRepo(s.DB)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "articles"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	_, err := repo.Count()
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_GetMostRead() {
	repo := NewArticleRepo(s.DB)
	today := time.Now()
	from := today.Add(time.Duration(-1) * time.Hour)
	to := today
	fromString := from.Format("2006-01-02 15:04:05")
	toString := to.Format("2006-01-02 15:04:05")

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id","title","description","link","published","authors",articles.articles_source_id,articles.created_at FROM "articles" JOIN (SELECT "article_id",count(article_id) as read FROM "reads" WHERE created_at between $1 AND $2 GROUP BY "article_id" ORDER BY read desc LIMIT 1) r on articles.id = r.article_id WHERE "articles"."deleted_at" IS NULL ORDER BY r.read desc,"articles"."id" LIMIT 1`)).
		WithArgs(fromString, toString).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description"}).
			AddRow("title test1", "description test1"))

	_, err := repo.GetMostRead(from, to)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestArticles_CreateIfNotExist() {
	repo := NewArticleRepo(s.DB)

	article := entities.Article{
		Title: "title test1",
		Description: "description test1",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "articles" WHERE "articles"."title" = $1 AND "articles"."deleted_at" IS NULL ORDER BY "articles"."id" LIMIT 1`)).
		WithArgs(article.Title).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description"}).
			AddRow("title test1", "description test1"))

	err := repo.CreateIfNotExist(&article)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}
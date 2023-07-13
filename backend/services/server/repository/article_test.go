package repository

// import (
// 	"database/sql"
// 	"regexp"

// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// 	"github.com/stretchr/testify/suite"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type Suite struct {
// 	suite.Suite
// 	DB   *gorm.DB
// 	mock sqlmock.Sqlmock
// }

// func (s *Suite) SetupSuite() {
// 	var (
// 		db  *sql.DB
// 		err error
// 	)

// 	db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
// 	require.NoError(s.T(), err)

// 	conn, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
// 	s.DB = conn
// 	require.NoError(s.T(), err)
// }

// func (s *Suite) AfterTest(_, _ string) {
// 	require.NoError(s.T(), s.mock.ExpectationsWereMet())
// }

// func TestInit(t *testing.T) {
// 	suite.Run(t, new(Suite))
// }

// func (s *Suite) TestArticles_GetArticlesPaginationByArticlesSourceID(t *testing.T) {
// 	s.mock.ExpectQuery(
// 		"SELECT `username`,`password` FROM `users` WHERE `username` = ?").
// 		WithArgs(want.Username).
// 		WillReturnRows(sqlmock.NewRows([]string{"username", "password"}).
// 			AddRow(want.Username, want.Password))
	
// 		articleRepo := NewArticleRepo(s.DB)
// 	page := 1;
// 	pageSize := 10;
// 	articlesSourceID := uint(1)
// 	articles, found, err := articleRepo.GetArticlesPaginationByArticlesSourceID(articlesSourceID, page, page)
// 	assert.Nil(s.T(), err)

// 	if *got != want {
// 		s.T().Error("Query result is wrong")
// 	}
// }
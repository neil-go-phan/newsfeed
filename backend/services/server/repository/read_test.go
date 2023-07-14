package repository

import (
	"regexp"
	"server/entities"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func (s *Suite) TestRead_Create() {
	repo := NewRead(s.DB)

	read := entities.Read{}

	s.mock.ExpectBegin()

	s.mock.ExpectExec("INSERT").
		WillReturnResult(sqlmock.NewResult(0, 1))

	s.mock.ExpectCommit()
	err := repo.Create(read)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestRead_Delete() {
	repo := NewRead(s.DB)

	read := entities.Read{
		Username: TEST_USERNAME,
		ArticleID: TEST_SOURCE_ID,
		ArticlesSourceID: TEST_ARTICLES_SOURCE_ID,
	}

	s.mock.ExpectBegin()

	s.mock.ExpectExec("DELETE").
		WillReturnResult(sqlmock.NewResult(0, 1))

	s.mock.ExpectCommit()
	err := repo.Delete(read)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}


func (s *Suite) TestRead_CountByUsernameAndSourceID() {
	repo := NewRead(s.DB)

	read := entities.Read{
		Username:         TEST_USERNAME,
		ArticlesSourceID: TEST_ARTICLES_SOURCE_ID,
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "reads" WHERE "reads"."username" = $1 AND "reads"."articles_source_id" = $2`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	_, err := repo.CountByUsernameAndSourceID(read)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestRead_SelectByUsernameOnDay() {
	repo := NewRead(s.DB)

	day := time.Now()
	dayString := day.Format("2006-01-02")
	from := dayString + " 00:00:00"
	to := dayString + " 23:59:59"

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "reads" WHERE username = $1 AND created_at BETWEEN $2 AND $3`)).
		WithArgs(TEST_USERNAME, from, to).
		WillReturnRows(sqlmock.NewRows([]string{"articles_source_id"}).
			AddRow(1))

	_, err := repo.SelectByUsernameOnDay(TEST_USERNAME, day)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestRead_SelectByUsernameAndSourceIDOnDay() {
	repo := NewRead(s.DB)

	day := time.Now()
	dayString := day.Format("2006-01-02")
	from := dayString + " 00:00:00"
	to := dayString + " 23:59:59"

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "reads" WHERE articles_source_id = $1 username = $2 AND created_at BETWEEN $3 AND $4`)).
		WithArgs(TEST_ARTICLES_SOURCE_ID, TEST_USERNAME, from, to).
		WillReturnRows(sqlmock.NewRows([]string{"articles_source_id"}).
			AddRow(1))

	_, err := repo.SelectByUsernameAndSourceIDOnDay(TEST_USERNAME, TEST_ARTICLES_SOURCE_ID, day)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestRead_MarkAllAsReadBySourceID() {
	repo := NewRead(s.DB)

	username := "test_user"
	articlesSourceID := uint(1)

	s.mock.ExpectExec("INSERT INTO reads").
		WithArgs(username, username, articlesSourceID, articlesSourceID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.MarkAllAsReadBySourceID(username, articlesSourceID)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestRead_MarkAllAsReadByUserFollowedSource() {
	repo := NewRead(s.DB)

	username := "test_user"

	s.mock.ExpectExec("INSERT INTO reads").
		WithArgs(username, username, username).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.MarkAllAsReadByUserFollowedSource(username)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

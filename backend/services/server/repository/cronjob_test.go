package repository

import (
	"regexp"
	"server/entities"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func (s *Suite) TestCronjob_Create() {
	repo := NewCronjobRepo(s.DB)

	cronjob := &entities.Cronjob{
		Name: "Hello world",
	}

	s.mock.ExpectBegin()
	s.mock.ExpectQuery("INSERT").
		WillReturnRows(sqlmock.NewRows([]string{"name", "illustration"}).
			AddRow("heelp", "base64"))
	s.mock.ExpectCommit()

	_, err := repo.Create(cronjob)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestCronjob_Get() {
	repo := NewCronjobRepo(s.DB)
	today := time.Now()
	from := today.Add(time.Duration(-1) *time.Hour)
	to := today
	fromString := from.Format("2006-01-02 15:04:05")
	toString := to.Format("2006-01-02 15:04:05")

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "cronjobs" WHERE (started_at BETWEEN $1 AND $2) AND "cronjobs"."deleted_at" IS NULL`)).
		WithArgs(fromString, toString).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description"}).
			AddRow("title test1", "description test1"))

	_, err := repo.Get(from, to)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestCronjob_GetRuning() {
	repo := NewCronjobRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "cronjobs" WHERE ended_at IS NULL AND "cronjobs"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"title", "description"}).
			AddRow("title test1", "description test1"))

	_,err := repo.GetRuning()
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}
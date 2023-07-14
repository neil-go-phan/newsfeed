package repository

import (
	"server/entities"

	"github.com/DATA-DOG/go-sqlmock"
)

func (s *Suite) TestReadLater_Create() {
	repo := NewReadLater(s.DB)

	readLater := entities.ReadLater{}

	s.mock.ExpectBegin()

	s.mock.ExpectExec("INSERT").
		WillReturnResult(sqlmock.NewResult(0, 1))

	s.mock.ExpectCommit()
	err := repo.Create(readLater)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestReadLater_DELETE() {
	repo := NewReadLater(s.DB)

	readLater := entities.ReadLater{
		Username:  TEST_USERNAME,
		ArticleID: TEST_SOURCE_ID,
	}

	s.mock.ExpectBegin()

	s.mock.ExpectExec("DELETE").
	WillReturnResult(sqlmock.NewResult(0, 1))

	s.mock.ExpectCommit()
	err := repo.Delete(readLater)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

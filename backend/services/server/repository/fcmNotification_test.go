package repository

import (
	"regexp"
	"server/entities"

	"github.com/DATA-DOG/go-sqlmock"
)

const FIREBASE_TOKEN = "token"

func (s *Suite) TestFcmNotification_Create() {
	repo := NewFcmNotification(s.DB)

	notify := entities.FcmNotification{}

	s.mock.ExpectBegin()

	s.mock.ExpectExec("INSERT").
		WillReturnResult(sqlmock.NewResult(0, 1))

	s.mock.ExpectCommit()
	err := repo.Create(notify)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestFcmNotification_List() {
	repo := NewFcmNotification(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "fcm_notifications"`)).
		WillReturnRows(sqlmock.NewRows([]string{"name", "illustration"}).
			AddRow("heelp", "base64"))

	_,err := repo.List()
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

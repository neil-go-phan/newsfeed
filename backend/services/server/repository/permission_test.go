package repository

import (
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
)

func (s *Suite) TestPermission_List() {
	repo := NewPermission(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "permissions" WHERE "permissions"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "illustration"}).
			AddRow(1, "name", "base64"))

	_, err := repo.List()
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}
package repository

import (
	"regexp"
	"server/entities"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

const TEST_ROLE_NAME = "Superadmin"

func (s *Suite) TestRole_Get() {
	repo := NewRoleRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "roles" WHERE name = $1 AND "roles"."deleted_at" IS NULL`)).
		WithArgs(TEST_ROLE_NAME).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).
			AddRow(TEST_ROLE_NAME))

	_, err := repo.Get(TEST_ROLE_NAME)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestRole_List() {
	repo := NewRoleRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "roles" WHERE "roles"."deleted_at" IS NULL LIMIT 10 OFFSET 10`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).
			AddRow(TEST_ROLE_NAME))

	_, err := repo.List(PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestRole_ListRoleName() {
	repo := NewRoleRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT "name" FROM "roles" WHERE "roles"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).
			AddRow(TEST_ROLE_NAME))

	_, err := repo.ListRoleName()
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestRole_Count() {
	repo := NewRoleRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "roles"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	_, err := repo.Count()
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestRole_Delete() {
	repo := NewRoleRepo(s.DB)

	s.mock.ExpectBegin()
	s.mock.ExpectExec("DELETE").
		WillReturnResult(sqlmock.NewResult(0, 1))

	s.mock.ExpectCommit()

	err := repo.Delete(TEST_USER_ID)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestRole_Update() {
	repo := NewRoleRepo(s.DB)

	role := entities.Role{
		Model: gorm.Model{
			ID: uint(1),
		},
		Name: "name",
		Description: "test",
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec("UPDATE").
		WillReturnResult(sqlmock.NewResult(0, 1))
	s.mock.ExpectCommit()

	s.mock.ExpectBegin()
	s.mock.ExpectExec("UPDATE").
	WillReturnResult(sqlmock.NewResult(0, 1))
	s.mock.ExpectCommit()

	s.mock.ExpectBegin()
	s.mock.ExpectExec("DELETE").
	WillReturnResult(sqlmock.NewResult(0, 1))
	s.mock.ExpectCommit()

	err := repo.Update(role)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

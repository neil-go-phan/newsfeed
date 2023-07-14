package repository

import (
	"regexp"
	"server/entities"

	"github.com/DATA-DOG/go-sqlmock"
)

const TEST_USER_ID = uint(1)

func (s *Suite) TestUser_Create() {
	repo := NewUserRepo(s.DB)

	userInput := &entities.User{}

	s.mock.ExpectBegin()

	s.mock.ExpectQuery("INSERT").
		WillReturnRows(sqlmock.NewRows([]string{"username", "email"}).
			AddRow("heelp", "base64"))

	s.mock.ExpectCommit()
	_, err := repo.Create(userInput)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestUser_ChangeRole() {
	repo := NewUserRepo(s.DB)

	s.mock.ExpectBegin()

	s.mock.ExpectExec("UPDATE").
		WillReturnResult(sqlmock.NewResult(0, 1))

	s.mock.ExpectCommit()
	err := repo.ChangeRole(TEST_USER_ID, TEST_ROLE_NAME)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestUser_UserUpgrateRole() {
	repo := NewUserRepo(s.DB)

	s.mock.ExpectBegin()

	s.mock.ExpectExec("UPDATE").
		WillReturnResult(sqlmock.NewResult(0, 1))

	s.mock.ExpectCommit()
	err := repo.UserUpgrateRole(TEST_ROLE_NAME)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestUser_Delete() {
	repo := NewUserRepo(s.DB)

	s.mock.ExpectBegin()

	s.mock.ExpectExec("DELETE").
		WillReturnResult(sqlmock.NewResult(0, 1))

	s.mock.ExpectCommit()
	err := repo.Delete(TEST_USER_ID)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestUser_List() {
	repo := NewUserRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL LIMIT 10 OFFSET 10`)).
		WillReturnRows(sqlmock.NewRows([]string{"username", "email"}).
			AddRow("title test1", "description test1"))

	_, err := repo.List(PAGE, PAGE_SIZE)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestUser_Count() {
	repo := NewUserRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "users"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	_, err := repo.Count()
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestUser_Get() {
	repo := NewUserRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT "role_name","email","username","password","salt" FROM "users" WHERE "username" = $1 AND "users"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"username", "email"}).
			AddRow("title test1", "description test1"))

	_, err := repo.Get(TEST_USERNAME)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestUser_GetWithEmail() {
	repo := NewUserRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "email" = $1 AND "users"."deleted_at" IS NULL`)).
		WithArgs(TEST_USERNAME).
		WillReturnRows(sqlmock.NewRows([]string{"username", "email"}).
			AddRow("title test1", "description test1"))

	_, err := repo.GetWithEmail(TEST_USERNAME)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestUser_FindOrCreateWithEmail() {
	repo := NewUserRepo(s.DB)

	user := entities.User{}

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`)).
		WillReturnRows(sqlmock.NewRows([]string{"username", "email"}).
			AddRow("heelp", "base64"))

	_, err := repo.FindOrCreateWithEmail(&user)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

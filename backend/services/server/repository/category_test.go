package repository

import (
	"regexp"
	"server/entities"

	"github.com/DATA-DOG/go-sqlmock"
)

func (s *Suite) TestCategory_CreateIfNotExist() {
	repo := NewCategory(s.DB)

	category := entities.Category{
		Name:         "heelp",
		Illustration: "base64",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories" WHERE "categories"."name" = $1 AND "categories"."deleted_at" IS NULL ORDER BY "categories"."id" LIMIT 1`)).
		WillReturnRows(sqlmock.NewRows([]string{"name", "illustration"}).
			AddRow("heelp", "base64"))

	err := repo.CreateIfNotExist(category)
	s.Assert().EqualError(err, "category already exist")
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestCategory_Get() {
	repo := NewCategory(s.DB)

	category := entities.Category{
		Name:         "heelp",
		Illustration: "base64",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories" WHERE name = $1 AND "categories"."deleted_at" IS NULL ORDER BY "categories"."id" LIMIT 1`)).
		WithArgs(category.Name).
		WillReturnRows(sqlmock.NewRows([]string{"name", "illustration"}).
			AddRow("heelp", "base64"))

	_, err := repo.Get(category.Name)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestCategory_ListName() {
	repo := NewCategory(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id","name" FROM "categories" WHERE "categories"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "base64"))

	_, err := repo.ListName()
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestCategory_ListAll() {
	repo := NewCategory(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id","name","illustration" FROM "categories" WHERE "categories"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "illustration"}).
			AddRow(1, "name", "base64"))

	_, err := repo.ListAll()
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestCategory_GetPagination() {
	repo := NewCategory(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories" WHERE "categories"."deleted_at" IS NULL LIMIT 10 OFFSET 10`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "illustration"}).
			AddRow(1, "name", "base64"))

	_, err := repo.GetPagination(PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestCategory_Count() {
	repo := NewCategory(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "categories"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	_, err := repo.Count()
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

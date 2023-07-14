package repository

import (
	"fmt"
	"regexp"
	"server/entities"
	"strings"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

const TEST_TOPIC_NAME = "topic"
const TEST_CATEGORY_ID = uint(1)

func (s *Suite) TestTopic_CreateIfNotExist() {
	repo := NewTopic(s.DB)

	topic := entities.Topic{
		Name: TEST_TOPIC_NAME,
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "topics" WHERE "topics"."name" = $1 AND "topics"."deleted_at" IS NULL ORDER BY "topics"."id" LIMIT 1`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).
			AddRow(TEST_TOPIC_NAME))

	err := repo.CreateIfNotExist(topic)
	s.Assert().EqualError(err, "topic already exist")
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestTopic_List() {
	repo := NewTopic(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "topics" WHERE "topics"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).
			AddRow(TEST_TOPIC_NAME))

	_, err := repo.List()
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestTopic_Count() {
	repo := NewTopic(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "topics"`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	_, err := repo.Count()
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestTopic_GetPagination() {
	repo := NewTopic(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "topics" WHERE "topics"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).
			AddRow(TEST_TOPIC_NAME))

	_, err := repo.GetPagination(PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestTopic_GetByCategory() {
	repo := NewTopic(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "topics" WHERE category_id = $1 AND "topics"."deleted_at" IS NUL`)).
		WithArgs(TEST_CATEGORY_ID).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).
			AddRow(TEST_TOPIC_NAME))

	_, err := repo.GetByCategory(TEST_CATEGORY_ID)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestTopic_SearchByName() {
	repo := NewTopic(s.DB)

	keyword := "keyword"
	searchKeyword := fmt.Sprint(strings.ToLower(keyword) + "%")

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "topics" WHERE LOWER(name) LIKE $1 AND "topics"."deleted_at" IS NULL`)).
		WithArgs(searchKeyword).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).
			AddRow(TEST_TOPIC_NAME))

	_, err := repo.SearchByName(keyword)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestTopic_Delete() {
	repo := NewTopic(s.DB)

	topic := entities.Topic{}

	s.mock.ExpectBegin()

	s.mock.ExpectExec("DELETE").
		WillReturnResult(sqlmock.NewResult(0, 1))

	s.mock.ExpectCommit()
	err := repo.Delete(topic)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestTopic_UpdateWhenDeteleCategory() {
	repo := NewTopic(s.DB)

	s.mock.ExpectBegin()
	s.mock.ExpectExec("UPDATE").
		WillReturnResult(sqlmock.NewResult(0, 1))

	s.mock.ExpectCommit()

	err := repo.UpdateWhenDeteleCategory(TEST_CATEGORY_ID, TEST_CATEGORY_ID)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestTopic_Update() {
	repo := NewTopic(s.DB)

	topic := entities.Topic{
		Model: gorm.Model{
			ID: TEST_TOPIC_ID,
		},
		Name: "name",
		CategoryID: TEST_CATEGORY_ID,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec("UPDATE").
	WillReturnResult(sqlmock.NewResult(0, 1))

	s.mock.ExpectCommit()

	err := repo.Update(topic)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

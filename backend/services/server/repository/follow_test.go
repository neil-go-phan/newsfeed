package repository

import (
	"regexp"
	"server/entities"

	"github.com/DATA-DOG/go-sqlmock"
)

func (s *Suite) TestFollow_CreateIfNotExist() {
	repo := NewFollow(s.DB)

	follow := entities.Follow{
		Username:         TEST_USERNAME,
		ArticlesSourceID: TEST_ARTICLES_SOURCE_ID,
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "follows" WHERE "follows"."username" = $1 AND "follows"."articles_source_id" = $2 ORDER BY "follows"."created_at" LIMIT 1`)).
		WithArgs(follow.Username, follow.ArticlesSourceID).
		WillReturnRows(sqlmock.NewRows([]string{"name", "illustration"}).
			AddRow("heelp", "base64"))

	err := repo.CreateIfNotExist(follow)
	s.Assert().EqualError(err, "user already follow this articles source")
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestFollow_GetByUsername() {
	repo := NewFollow(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "follows" WHERE "follows"."username" = $1`)).
		WithArgs(TEST_USERNAME).
		WillReturnRows(sqlmock.NewRows([]string{"name", "illustration"}).
			AddRow("heelp", "base64"))

	_, err := repo.GetByUsername(TEST_USERNAME)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestFollow_GetNewestFeedsUpdated() {
	repo := NewFollow(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT "articles_source_id" FROM "follows" WHERE username = $1 AND unread > 0 ORDER BY updated_at desc LIMIT 3`)).
		WithArgs(TEST_USERNAME).
		WillReturnRows(sqlmock.NewRows([]string{"name", "illustration"}).
			AddRow("heelp", "base64"))

	_, err := repo.GetNewestFeedsUpdated(TEST_USERNAME)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestFollow_Delete() {
	repo := NewFollow(s.DB)

	follow := entities.Follow{
		Username: TEST_USERNAME,
		ArticlesSourceID: TEST_ARTICLES_SOURCE_ID,
	}

	s.mock.ExpectBegin()

	s.mock.ExpectExec("DELETE").
		WillReturnResult(sqlmock.NewResult(0, 1))

	s.mock.ExpectCommit()
	err := repo.Delete(follow)
	s.Assert().NoError(err)

	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}
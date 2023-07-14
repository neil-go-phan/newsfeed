package repository

import (
	"regexp"
	"server/entities"

	"github.com/DATA-DOG/go-sqlmock"
)

const TEST_CRAWLER_ID = 1
const TEST_SOURCE_LINK = "https://test.com"

func (s *Suite) TestCrawler_Get() {
	repo := NewCrawlerRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "crawlers" WHERE "crawlers"."deleted_at" IS NULL AND "crawlers"."id" = $1 ORDER BY "crawlers"."id" LIMIT 1`)).
		WithArgs(TEST_CRAWLER_ID).
		WillReturnRows(sqlmock.NewRows([]string{"name", "illustration"}).
			AddRow("heelp", "base64"))

	_, err := repo.Get(TEST_CRAWLER_ID)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestCrawler_GetBySourceLink() {
	repo := NewCrawlerRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "crawlers" WHERE source_link = $1 AND "crawlers"."deleted_at" IS NULL ORDER BY "crawlers"."id" LIMIT 1`)).
		WithArgs(TEST_SOURCE_LINK).
		WillReturnRows(sqlmock.NewRows([]string{"name", "illustration"}).
			AddRow("heelp", "base64"))

	_, err := repo.GetBySourceLink(TEST_SOURCE_LINK)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestCrawler_CreateIfNotExist() {
	repo := NewCrawlerRepo(s.DB)

	crawler := entities.Crawler{
		SourceLink: TEST_SOURCE_LINK,
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "crawlers" WHERE "crawlers"."source_link" = $1 AND "crawlers"."deleted_at" IS NULL ORDER BY "crawlers"."id" LIMIT 1`)).
		WillReturnRows(sqlmock.NewRows([]string{"name", "illustration"}).
			AddRow("heelp", "base64"))

	_, err := repo.CreateIfNotExist(crawler)
	s.Assert().EqualError(err, "crawler already exist")
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestCrawler_List() {
	repo := NewCrawlerRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "crawlers" WHERE "crawlers"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"name", "illustration"}).
			AddRow("heelp", "base64"))

	_, err := repo.List()
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

func (s *Suite) TestCrawler_ListAllPaging() {
	repo := NewCrawlerRepo(s.DB)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "crawlers" WHERE "crawlers"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"name", "illustration"}).
			AddRow("heelp", "base64"))

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "crawlers" WHERE "crawlers"."deleted_at" IS NULL LIMIT 10 OFFSET 10`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).
			AddRow(1))

	_, _, err := repo.ListAllPaging(PAGE, PAGE_SIZE)
	s.Assert().NoError(err)
	err = s.mock.ExpectationsWereMet()
	s.Assert().NoError(err)
}

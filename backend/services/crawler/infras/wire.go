// go:build wireinject
// + build wireinject
// go:generate go run github.com/google/wire/cmd/wire@latest
package infras

import (
	"crawler/handlers"
	"crawler/repository"
	"crawler/services"
	"crawler/services/articlesSource"
	"crawler/services/article"
	"crawler/services/crawler"
	"crawler/services/follow"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitGRPCServer(db *gorm.DB) *handlers.GRPCServer {
	wire.Build(
		repository.NewArticleRepo,
		wire.Bind(new(repository.ArticleRepository), new(*repository.ArticleRepo)),
		articleservices.NewArticleService,
		wire.Bind(new(services.ArticleServices), new(*articleservices.ArticleService)),

		repository.NewArticlesSourcesRepo,
		wire.Bind(new(repository.ArticlesSourcesRepository), new(*repository.ArticlesSourcesRepo)),
		articlessourceservices.NewArticlesSourceService,
		wire.Bind(new(services.ArticlesSourceServices), new(*articlessourceservices.ArticlesSourceService)),

		repository.NewFollow,
		wire.Bind(new(repository.FollowRepository), new(*repository.FollowRepo)),
		followservices.NewFollowService,
		wire.Bind(new(services.FollowServices), new(*followservices.FollowService)),

		repository.NewCrawlerRepo,
		wire.Bind(new(repository.CrawlerRepository), new(*repository.CrawlerRepo)),
		crawlerservices.NewCrawlerService,
		wire.Bind(new(services.CrawlerServices), new(*crawlerservices.CrawlerService)),

		handlers.NewGRPCServer,
	)
	return &handlers.GRPCServer{}
}
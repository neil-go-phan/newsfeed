// go:build wireinject
// + build wireinject
// go:generate go run github.com/google/wire/cmd/wire@latest
package infras

import (
	"server/handlers"
	pb "server/proto"
	"server/repository"
	"server/routes"
	"server/services"
	"server/services/article"
	"server/services/articlesSource"
	"server/services/crawler"
	"server/services/cronjob"
	"server/services/role"
	"server/services/user"

	"github.com/google/wire"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func InitUser(db *gorm.DB) *routes.UserRoutes {
	wire.Build(
		repository.NewUserRepo,
		repository.NewRoleRepo,
		userservice.NewUserService,
		roleservice.NewRoleService,
		handlers.NewUserHandler,
		routes.NewUserRoutes,
		wire.Bind(new(repository.UserRepository), new(*repository.UserRepo)),
		wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepo)),
		wire.Bind(new(services.UserServices), new(*userservice.UserService)),
		wire.Bind(new(services.RoleServices), new(*roleservice.RoleService)),
		wire.Bind(new(handlers.UserHandlerInterface), new(*handlers.UserHandler)),
	)
	return &routes.UserRoutes{}
}

func InitCrawler(db *gorm.DB, grpcClient pb.CrawlerServiceClient, cronjob *cron.Cron, jobIDMap map[string]cron.EntryID) *routes.CrawlerRoutes {
	wire.Build(
		repository.NewArticleRepo,
		wire.Bind(new(repository.ArticleRepository), new(*repository.ArticleRepo)),
		articleservices.NewArticleService,
		wire.Bind(new(services.ArticleServices), new(*articleservices.ArticleService)),

		repository.NewArticlesSourcesRepo,
		wire.Bind(new(repository.ArticlesSourcesRepository), new(*repository.ArticlesSourcesRepo)),
		articlessourceservices.NewArticlesSourceService,
		wire.Bind(new(services.ArticlesSourceServices), new(*articlessourceservices.ArticlesSourceService)),

		repository.NewCronjobRepo,
		wire.Bind(new(repository.CronjobRepository), new(*repository.CronjobRepo)),
		cronjobservices.NewCronjobService,
		wire.Bind(new(services.CronjobServices), new(*cronjobservices.CronjobService)),

		repository.NewCrawlerRepo,
		wire.Bind(new(repository.CrawlerRepository), new(*repository.CrawlerRepo)),
		crawlerservices.NewCrawlerService,
		wire.Bind(new(services.CrawlerServices), new(*crawlerservices.CrawlerService)),
		handlers.NewCrawlerHandler,
		wire.Bind(new(handlers.CrawlerHandlerInterface), new(*handlers.CrawlerHandler)),
		routes.NewCrawlerRoutes,
	)
	return &routes.CrawlerRoutes{}
}

// func InitCronjob(db *gorm.DB, grpcClient pb.CrawlerServiceClient) *routes.CrawlerRoutes {
// 	wire.Build(
// 		repository.NewArticleRepo,
// 		wire.Bind(new(repository.ArticleRepository), new(*repository.ArticleRepo)),
// 		articleservices.NewArticleService,
// 		wire.Bind(new(services.ArticleServices), new(*articleservices.ArticleService)),

// 		repository.NewArticlesSourcesRepo,
// 		wire.Bind(new(repository.ArticlesSourcesRepository), new(*repository.ArticlesSourcesRepo)),
// 		articlessourceservices.NewArticlesSourceService,
// 		wire.Bind(new(services.ArticlesSourceServices), new(*articlessourceservices.ArticlesSourceService)),

// 		repository.NewCrawlerRepo,
// 		wire.Bind(new(repository.CrawlerRepository), new(*repository.CrawlerRepo)),
// 		crawlerservices.NewCrawlerService,
// 		wire.Bind(new(services.CrawlerServices), new(*crawlerservices.CrawlerService)),
// 		handlers.NewCrawlerHandler,
// 		wire.Bind(new(handlers.CrawlerHandlerInterface), new(*handlers.CrawlerHandler)),
// 		routes.NewCrawlerRoutes,
// 	)
// 	return &routes.CrawlerRoutes{}
// }
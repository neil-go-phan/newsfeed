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
	articleservices "server/services/article"
	articlessourceservices "server/services/articlesSource"
	categoryservices "server/services/category"
	crawlerservices "server/services/crawler"
	cronjobservices "server/services/cronjob"
	followservices "server/services/follow"
	readservices "server/services/read"
	readlaterservices "server/services/readLater"
	roleservice "server/services/role"
	topicservices "server/services/topic"
	userservices "server/services/user"

	"server/services/fcmNotification"
	"server/services/permission"

	"firebase.google.com/go/messaging"
	"github.com/google/wire"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func InitUser(db *gorm.DB) *routes.UserRoutes {
	wire.Build(
		repository.NewUserRepo,
		repository.NewRoleRepo,
		userservices.NewUserService,
		roleservice.NewRoleService,
		handlers.NewUserHandler,
		routes.NewUserRoutes,
		wire.Bind(new(repository.UserRepository), new(*repository.UserRepo)),
		wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepo)),
		wire.Bind(new(services.UserServices), new(*userservices.UserService)),
		wire.Bind(new(services.RoleServices), new(*roleservice.RoleService)),
		wire.Bind(new(handlers.UserHandlerInterface), new(*handlers.UserHandler)),
	)
	return &routes.UserRoutes{}
}

func InitCrawler(db *gorm.DB, grpcClient pb.CrawlerServiceClient, cronjob *cron.Cron, jobIDMap map[string]cron.EntryID) *routes.CrawlerRoutes {
	wire.Build(
		repository.NewRoleRepo,
		wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepo)),
		roleservice.NewRoleService,
		wire.Bind(new(services.RoleServices), new(*roleservice.RoleService)),

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

func InitTopic(db *gorm.DB) *routes.TopicRoutes {
	wire.Build(
		repository.NewRoleRepo,
		wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepo)),
		roleservice.NewRoleService,
		wire.Bind(new(services.RoleServices), new(*roleservice.RoleService)),

		repository.NewArticlesSourcesRepo,
		wire.Bind(new(repository.ArticlesSourcesRepository), new(*repository.ArticlesSourcesRepo)),
		articlessourceservices.NewArticlesSourceService,
		wire.Bind(new(services.ArticlesSourceServices), new(*articlessourceservices.ArticlesSourceService)),

		repository.NewTopic,
		wire.Bind(new(repository.TopicRepository), new(*repository.TopicRepo)),
		topicservices.NewTopicService,
		wire.Bind(new(services.TopicServices), new(*topicservices.TopicService)),
		handlers.NewTopicHandler,
		wire.Bind(new(handlers.TopicHandlerInterface), new(*handlers.TopicHandler)),
		routes.NewTopicRoutes,
	)
	return &routes.TopicRoutes{}
}

func InitCategory(db *gorm.DB) *routes.CategoryRoutes {
	wire.Build(
		repository.NewRoleRepo,
		wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepo)),
		roleservice.NewRoleService,
		wire.Bind(new(services.RoleServices), new(*roleservice.RoleService)),

		repository.NewArticlesSourcesRepo,
		wire.Bind(new(repository.ArticlesSourcesRepository), new(*repository.ArticlesSourcesRepo)),
		articlessourceservices.NewArticlesSourceService,
		wire.Bind(new(services.ArticlesSourceServices), new(*articlessourceservices.ArticlesSourceService)),

		repository.NewTopic,
		wire.Bind(new(repository.TopicRepository), new(*repository.TopicRepo)),
		topicservices.NewTopicService,
		wire.Bind(new(services.TopicServices), new(*topicservices.TopicService)),

		repository.NewCategory,
		wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepo)),
		categoryservices.NewCategoryService,
		wire.Bind(new(services.CategoryServices), new(*categoryservices.CategoryService)),
		handlers.NewCategoryHandler,
		wire.Bind(new(handlers.CategoryHandlerInterface), new(*handlers.CategoryHandler)),
		routes.NewCategoryRoutes,
	)
	return &routes.CategoryRoutes{}
}

func InitArticlesSources(db *gorm.DB) *routes.ArticlesSourceRoutes {
	wire.Build(
		repository.NewRoleRepo,
		wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepo)),
		roleservice.NewRoleService,
		wire.Bind(new(services.RoleServices), new(*roleservice.RoleService)),

		repository.NewArticlesSourcesRepo,
		wire.Bind(new(repository.ArticlesSourcesRepository), new(*repository.ArticlesSourcesRepo)),
		articlessourceservices.NewArticlesSourceService,
		wire.Bind(new(services.ArticlesSourceServices), new(*articlessourceservices.ArticlesSourceService)),
		handlers.NewArticlesSourceHandler,
		wire.Bind(new(handlers.ArticlesSourceHandlerInterface), new(*handlers.ArticlesSourceHandler)),
		routes.NewArticlesSourceRoutes,
	)
	return &routes.ArticlesSourceRoutes{}
}

func InitArticles(db *gorm.DB) *routes.ArticleRoutes {
	wire.Build(
		repository.NewRoleRepo,
		wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepo)),
		roleservice.NewRoleService,
		wire.Bind(new(services.RoleServices), new(*roleservice.RoleService)),

		repository.NewArticleRepo,
		wire.Bind(new(repository.ArticleRepository), new(*repository.ArticleRepo)),
		articleservices.NewArticleService,
		wire.Bind(new(services.ArticleServices), new(*articleservices.ArticleService)),
		handlers.NewArticlesHandler,
		wire.Bind(new(handlers.ArticleHandlerInterface), new(*handlers.ArticleHandler)),
		routes.NewArticleRoutes,
	)
	return &routes.ArticleRoutes{}
}

func InitFollow(db *gorm.DB) *routes.FollowRoutes {
	wire.Build(
		repository.NewRoleRepo,
		wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepo)),
		roleservice.NewRoleService,
		wire.Bind(new(services.RoleServices), new(*roleservice.RoleService)),

		repository.NewArticlesSourcesRepo,
		wire.Bind(new(repository.ArticlesSourcesRepository), new(*repository.ArticlesSourcesRepo)),
		articlessourceservices.NewArticlesSourceService,
		wire.Bind(new(services.ArticlesSourceServices), new(*articlessourceservices.ArticlesSourceService)),

		repository.NewFollow,
		wire.Bind(new(repository.FollowRepository), new(*repository.FollowRepo)),
		followservices.NewFollowService,
		wire.Bind(new(services.FollowServices), new(*followservices.FollowService)),
		handlers.NewFollowHandler,
		wire.Bind(new(handlers.FollowHandlerInterface), new(*handlers.FollowHandler)),
		routes.NewFollowRoutes,
	)
	return &routes.FollowRoutes{}
}

func InitRead(db *gorm.DB) *routes.ReadRoutes {
	wire.Build(
		repository.NewRead,
		wire.Bind(new(repository.ReadRepository), new(*repository.ReadRepo)),
		readservices.NewReadService,
		wire.Bind(new(services.ReadServices), new(*readservices.ReadService)),
		handlers.NewReadHandler,
		wire.Bind(new(handlers.ReadHandlerInterface), new(*handlers.ReadHandler)),
		routes.NewReadRoutes,
	)
	return &routes.ReadRoutes{}
}

func InitRole(db *gorm.DB) *routes.RoleRoutes {
	wire.Build(
		repository.NewRoleRepo,
		wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepo)),
		roleservice.NewRoleService,
		wire.Bind(new(services.RoleServices), new(*roleservice.RoleService)),
		handlers.NewRoleHandler,
		wire.Bind(new(handlers.RoleHandlerInterface), new(*handlers.RoleHandler)),
		routes.NewRoleRoutes,
	)
	return &routes.RoleRoutes{}
}

func InitPermission(db *gorm.DB) *routes.PermissionRoutes {
	wire.Build(
		repository.NewPermission,
		wire.Bind(new(repository.PermissionRepository), new(*repository.PermissionRepo)),
		permissionservice.NewPermissionService,
		wire.Bind(new(services.PermissionServices), new(*permissionservice.PermissionService)),
		handlers.NewPermissionHandler,
		wire.Bind(new(handlers.PermissionHandlerInterface), new(*handlers.PermissionHandler)),
		routes.NewPermissionRoutes,
	)
	return &routes.PermissionRoutes{}
}

func InitReadLater(db *gorm.DB) *routes.ReadLaterRoutes {
	wire.Build(
		repository.NewRoleRepo,
		wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepo)),
		roleservice.NewRoleService,
		wire.Bind(new(services.RoleServices), new(*roleservice.RoleService)),

		repository.NewReadLater,
		wire.Bind(new(repository.ReadLaterRepository), new(*repository.ReadLaterRepo)),
		readlaterservices.NewReadLaterService,
		wire.Bind(new(services.ReadLaterServices), new(*readlaterservices.ReadLaterService)),
		handlers.NewReadLaterHandler,
		wire.Bind(new(handlers.ReadLaterHandlerInterface), new(*handlers.ReadLaterHandler)),
		routes.NewReadLaterRoutes,
	)
	return &routes.ReadLaterRoutes{}
}

func InitFcmNotification(db *gorm.DB,fcmClient *messaging.Client, cron *cron.Cron) *routes.FcmNotificationRoutes {
	wire.Build(
		repository.NewRoleRepo,
		wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepo)),
		roleservice.NewRoleService,
		wire.Bind(new(services.RoleServices), new(*roleservice.RoleService)),

		repository.NewArticleRepo,
		wire.Bind(new(repository.ArticleRepository), new(*repository.ArticleRepo)),
		articleservices.NewArticleService,
		wire.Bind(new(services.ArticleServices), new(*articleservices.ArticleService)),

		repository.NewFcmNotification,
		wire.Bind(new(repository.FcmNotificationRepository), new(*repository.FcmNotificationRepo)),
		notificationservices.NewFcmNotificationService,
		wire.Bind(new(services.FcmNotificationServices), new(*notificationservices.FcmNotificationService)),
		handlers.NewFcmNotificationHandler,
		wire.Bind(new(handlers.FcmNotificationHandlerInterface), new(*handlers.FcmNotificationHandler)),
		routes.NewFcmNotificationRoutes,
	)
	return &routes.FcmNotificationRoutes{}
}

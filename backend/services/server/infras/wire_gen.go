// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package infras

import (
	"firebase.google.com/go/messaging"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"server/handlers"
	"server/proto"
	"server/repository"
	"server/routes"
	"server/services/article"
	"server/services/articlesSource"
	"server/services/category"
	"server/services/crawler"
	"server/services/cronjob"
	"server/services/fcmNotification"
	"server/services/follow"
	"server/services/permission"
	"server/services/read"
	"server/services/readLater"
	"server/services/role"
	"server/services/topic"
	"server/services/user"
)

// Injectors from wire.go:

func InitizeUser(db *gorm.DB) *routes.UserRoutes {
	userRepo := repository.NewUserRepo(db)
	roleRepo := repository.NewRoleRepo(db)
	roleService := roleservice.NewRoleService(roleRepo)
	userService := userservice.NewUserService(userRepo, roleService)
	userHandler := handlers.NewUserHandler(userService)
	userRoutes := routes.NewUserRoutes(userHandler)
	return userRoutes
}

func InitizeCrawler(db *gorm.DB, grpcClient serverproto.CrawlerServiceClient, cronjob *cron.Cron, jobIDMap map[string]cron.EntryID) *routes.CrawlerRoutes {
	crawlerRepo := repository.NewCrawlerRepo(db)
	articleRepo := repository.NewArticleRepo(db)
	roleRepo := repository.NewRoleRepo(db)
	roleService := roleservice.NewRoleService(roleRepo)
	articleService := articleservices.NewArticleService(articleRepo, roleService)
	articlesSourcesRepo := repository.NewArticlesSourcesRepo(db)
	articlesSourceService := articlessourceservices.NewArticlesSourceService(articlesSourcesRepo, roleService)
	cronjobRepo := repository.NewCronjobRepo(db)
	cronjobService := cronjobservices.NewCronjobService(cronjobRepo, cronjob, grpcClient, jobIDMap)
	crawlerService := crawlerservices.NewCrawlerService(crawlerRepo, articleService, articlesSourceService, cronjobService, grpcClient, roleService)
	crawlerHandler := handlers.NewCrawlerHandler(crawlerService)
	crawlerRoutes := routes.NewCrawlerRoutes(crawlerHandler)
	return crawlerRoutes
}

func InitizeTopic(db *gorm.DB) *routes.TopicRoutes {
	topicRepo := repository.NewTopic(db)
	articlesSourcesRepo := repository.NewArticlesSourcesRepo(db)
	roleRepo := repository.NewRoleRepo(db)
	roleService := roleservice.NewRoleService(roleRepo)
	articlesSourceService := articlessourceservices.NewArticlesSourceService(articlesSourcesRepo, roleService)
	topicService := topicservices.NewTopicService(topicRepo, articlesSourceService, roleService)
	topicHandler := handlers.NewTopicHandler(topicService)
	topicRoutes := routes.NewTopicRoutes(topicHandler)
	return topicRoutes
}

func InitizeCategory(db *gorm.DB) *routes.CategoryRoutes {
	categoryRepo := repository.NewCategory(db)
	topicRepo := repository.NewTopic(db)
	articlesSourcesRepo := repository.NewArticlesSourcesRepo(db)
	roleRepo := repository.NewRoleRepo(db)
	roleService := roleservice.NewRoleService(roleRepo)
	articlesSourceService := articlessourceservices.NewArticlesSourceService(articlesSourcesRepo, roleService)
	topicService := topicservices.NewTopicService(topicRepo, articlesSourceService, roleService)
	categoryService := categoryservices.NewCategoryService(categoryRepo, topicService, roleService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	categoryRoutes := routes.NewCategoryRoutes(categoryHandler)
	return categoryRoutes
}

func InitizeArticlesSources(db *gorm.DB) *routes.ArticlesSourceRoutes {
	articlesSourcesRepo := repository.NewArticlesSourcesRepo(db)
	roleRepo := repository.NewRoleRepo(db)
	roleService := roleservice.NewRoleService(roleRepo)
	articlesSourceService := articlessourceservices.NewArticlesSourceService(articlesSourcesRepo, roleService)
	articlesSourceHandler := handlers.NewArticlesSourceHandler(articlesSourceService)
	articlesSourceRoutes := routes.NewArticlesSourceRoutes(articlesSourceHandler)
	return articlesSourceRoutes
}

func InitizeArticles(db *gorm.DB) *routes.ArticleRoutes {
	articleRepo := repository.NewArticleRepo(db)
	roleRepo := repository.NewRoleRepo(db)
	roleService := roleservice.NewRoleService(roleRepo)
	articleService := articleservices.NewArticleService(articleRepo, roleService)
	articleHandler := handlers.NewArticlesHandler(articleService)
	articleRoutes := routes.NewArticleRoutes(articleHandler)
	return articleRoutes
}

func InitizeFollow(db *gorm.DB) *routes.FollowRoutes {
	followRepo := repository.NewFollow(db)
	articlesSourcesRepo := repository.NewArticlesSourcesRepo(db)
	roleRepo := repository.NewRoleRepo(db)
	roleService := roleservice.NewRoleService(roleRepo)
	articlesSourceService := articlessourceservices.NewArticlesSourceService(articlesSourcesRepo, roleService)
	followService := followservices.NewFollowService(followRepo, articlesSourceService)
	followHandler := handlers.NewFollowHandler(followService)
	followRoutes := routes.NewFollowRoutes(followHandler)
	return followRoutes
}

func InitizeRead(db *gorm.DB) *routes.ReadRoutes {
	readRepo := repository.NewRead(db)
	readService := readservices.NewReadService(readRepo)
	readHandler := handlers.NewReadHandler(readService)
	readRoutes := routes.NewReadRoutes(readHandler)
	return readRoutes
}

func InitizeRole(db *gorm.DB) *routes.RoleRoutes {
	roleRepo := repository.NewRoleRepo(db)
	roleService := roleservice.NewRoleService(roleRepo)
	roleHandler := handlers.NewRoleHandler(roleService)
	roleRoutes := routes.NewRoleRoutes(roleHandler)
	return roleRoutes
}

func InitizePermission(db *gorm.DB) *routes.PermissionRoutes {
	permissionRepo := repository.NewPermission(db)
	permissionService := permissionservice.NewPermissionService(permissionRepo)
	permissionHandler := handlers.NewPermissionHandler(permissionService)
	permissionRoutes := routes.NewPermissionRoutes(permissionHandler)
	return permissionRoutes
}

func InitizeReadLater(db *gorm.DB) *routes.ReadLaterRoutes {
	readLaterRepo := repository.NewReadLater(db)
	roleRepo := repository.NewRoleRepo(db)
	roleService := roleservice.NewRoleService(roleRepo)
	readLaterService := readlaterservices.NewReadLaterService(readLaterRepo, roleService)
	readLaterHandler := handlers.NewReadLaterHandler(readLaterService)
	readLaterRoutes := routes.NewReadLaterRoutes(readLaterHandler)
	return readLaterRoutes
}

func InitizeFcmNotification(db *gorm.DB, fcmClient *messaging.Client, cron2 *cron.Cron) *routes.FcmNotificationRoutes {
	fcmNotificationRepo := repository.NewFcmNotification(db)
	articleRepo := repository.NewArticleRepo(db)
	roleRepo := repository.NewRoleRepo(db)
	roleService := roleservice.NewRoleService(roleRepo)
	articleService := articleservices.NewArticleService(articleRepo, roleService)
	fcmNotificationService := notificationservices.NewFcmNotificationService(fcmNotificationRepo, articleService, fcmClient, cron2)
	fcmNotificationHandler := handlers.NewFcmNotificationHandler(fcmNotificationService)
	fcmNotificationRoutes := routes.NewFcmNotificationRoutes(fcmNotificationHandler)
	return fcmNotificationRoutes
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package infras

import (
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
	articleService := articleservices.NewArticleService(articleRepo)
	articlesSourcesRepo := repository.NewArticlesSourcesRepo(db)
	articlesSourceService := articlessourceservices.NewArticlesSourceService(articlesSourcesRepo)
	cronjobRepo := repository.NewCronjobRepo(db)
	cronjobService := cronjobservices.NewCronjobService(cronjobRepo, cronjob, grpcClient, jobIDMap)
	crawlerService := crawlerservices.NewCrawlerService(crawlerRepo, articleService, articlesSourceService, cronjobService, grpcClient)
	crawlerHandler := handlers.NewCrawlerHandler(crawlerService)
	crawlerRoutes := routes.NewCrawlerRoutes(crawlerHandler)
	return crawlerRoutes
}

func InitizeTopic(db *gorm.DB) *routes.TopicRoutes {
	topicRepo := repository.NewTopic(db)
	articlesSourcesRepo := repository.NewArticlesSourcesRepo(db)
	articlesSourceService := articlessourceservices.NewArticlesSourceService(articlesSourcesRepo)
	topicService := topicservices.NewTopicService(topicRepo, articlesSourceService)
	topicHandler := handlers.NewTopicHandler(topicService)
	topicRoutes := routes.NewTopicRoutes(topicHandler)
	return topicRoutes
}

func InitizeCategory(db *gorm.DB) *routes.CategoryRoutes {
	categoryRepo := repository.NewCategory(db)
	topicRepo := repository.NewTopic(db)
	articlesSourcesRepo := repository.NewArticlesSourcesRepo(db)
	articlesSourceService := articlessourceservices.NewArticlesSourceService(articlesSourcesRepo)
	topicService := topicservices.NewTopicService(topicRepo, articlesSourceService)
	categoryService := categoryservices.NewCategoryService(categoryRepo, topicService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	categoryRoutes := routes.NewCategoryRoutes(categoryHandler)
	return categoryRoutes
}

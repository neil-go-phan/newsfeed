package routes

import (
	"server/handlers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

type CategoryRoutes struct {
	handler handlers.CategoryHandlerInterface
}

func NewCategoryRoutes(handler handlers.CategoryHandlerInterface) *CategoryRoutes {
	route := &CategoryRoutes{
		handler: handler,
	}
	return route
}

func (route *CategoryRoutes) Setup(r *gin.Engine) {
	routes := r.Group("category")
	{
		routes.GET("list/name", middlewares.CheckAccessToken(), route.handler.ListName)
		routes.GET("list/all", middlewares.CheckAccessToken(), route.handler.ListAll)
		routes.GET("get/page", middlewares.CheckAccessToken(), route.handler.GetPagination)
		routes.GET("count", middlewares.CheckAccessToken(), route.handler.Count)
		routes.POST("create", middlewares.CheckAccessToken(), route.handler.Create)
		routes.POST("update", middlewares.CheckAccessToken(), route.handler.Update)
		routes.POST("delete", middlewares.CheckAccessToken(), route.handler.Delete)
	}
}

package routes

import (
	"server/handlers"
	"server/middlewares"
	"github.com/gin-gonic/gin"
)

type CategoryRoutes struct {
	handler handlers.CategoryHandlerInterface
}

func NewCategoryRoutes(handler handlers.CategoryHandlerInterface) *CategoryRoutes{
	route := &CategoryRoutes{
		handler: handler,
	}
	return route
}

func (route *CategoryRoutes)Setup(r *gin.Engine) {
	routes := r.Group("category")
	{
		routes.GET("list", middlewares.CheckAccessToken(), route.handler.List)
		routes.GET("get-page", middlewares.CheckAccessToken(), route.handler.GetPagination)
		routes.GET("count", middlewares.CheckAccessToken(), route.handler.Count)
		routes.POST("create", middlewares.CheckAccessToken(), route.handler.Create)
		routes.POST("update-name", middlewares.CheckAccessToken(), route.handler.UpdateName)
		routes.POST("delete", middlewares.CheckAccessToken(), route.handler.Delete)
	}
}
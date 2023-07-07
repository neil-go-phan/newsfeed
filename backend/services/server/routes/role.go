package routes

import (
	"server/handlers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

type RoleRoutes struct {
	handler handlers.RoleHandlerInterface
}

func NewRoleRoutes(handler handlers.RoleHandlerInterface) *RoleRoutes {
	route := &RoleRoutes{
		handler: handler,
	}
	return route
}

func (route *RoleRoutes) Setup(r *gin.Engine) {
	routes := r.Group("role")
	{
		routes.GET("get", middlewares.CheckAccessToken(), route.handler.Get)
		routes.GET("list", middlewares.CheckAccessToken(), route.handler.List)
		routes.GET("list/names", middlewares.CheckAccessToken(), route.handler.ListNames)

		routes.GET("count/all", middlewares.CheckAccessToken(), route.handler.Total)
		routes.POST("create", middlewares.CheckAccessToken(), route.handler.Create)
		routes.POST("delete", middlewares.CheckAccessToken(), route.handler.Delete)
		routes.POST("update", middlewares.CheckAccessToken(), route.handler.Update)
	}
}

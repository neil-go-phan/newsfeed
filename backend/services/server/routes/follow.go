package routes

import (
	"server/handlers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

type FollowRoutes struct {
	handler handlers.FollowHandlerInterface
}

func NewFollowRoutes(handler handlers.FollowHandlerInterface) *FollowRoutes {
	route := &FollowRoutes{
		handler: handler,
	}
	return route
}

func (route *FollowRoutes) Setup(r *gin.Engine) {
	routes := r.Group("follow")
	{
		routes.GET("follow", middlewares.CheckAccessToken(), route.handler.Follow)
		routes.GET("unfollow", middlewares.CheckAccessToken(), route.handler.Unfollow)
		routes.GET("get/articles-sources", middlewares.CheckAccessToken(), route.handler.GetArticleSourceFollowed)
	}
}

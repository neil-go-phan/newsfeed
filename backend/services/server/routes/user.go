package routes

import (
	"server/handlers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userHandler handlers.UserHandlerInterface
}

func NewUserRoutes(userHandler handlers.UserHandlerInterface) *UserRoutes {
	userRoute := &UserRoutes{
		userHandler: userHandler,
	}
	return userRoute
}

func (userRoutes *UserRoutes) Setup(r *gin.Engine) {
	authRoutes := r.Group("auth")
	{
		authRoutes.GET("check-auth", middlewares.CheckAccessToken(), userRoutes.userHandler.CheckAuth)
		authRoutes.GET("token", middlewares.ExpiredAccessTokenHandler(), userRoutes.userHandler.Token)
		authRoutes.POST("register", userRoutes.userHandler.Register)
		authRoutes.POST("login", userRoutes.userHandler.Login)
		authRoutes.GET("oauth/google", userRoutes.userHandler.GoogleOAuth)
		authRoutes.GET("access-admin", middlewares.CheckAccessToken(), userRoutes.userHandler.AccessAdminPage)
		authRoutes.GET("list", middlewares.CheckAccessToken(), userRoutes.userHandler.List)
		authRoutes.POST("delete", middlewares.CheckAccessToken(), userRoutes.userHandler.Delete)
		authRoutes.POST("update/role", middlewares.CheckAccessToken(), userRoutes.userHandler.ChangeRole)
		authRoutes.GET("update/premium", middlewares.CheckAccessToken(), userRoutes.userHandler.UserUpgrateRole)

		authRoutes.GET("count", middlewares.CheckAccessToken(), userRoutes.userHandler.Total)

	}
}

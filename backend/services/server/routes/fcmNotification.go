package routes

import (
	"server/handlers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

type FcmNotificationRoutes struct {
	handler handlers.FcmNotificationHandlerInterface
}

func NewFcmNotificationRoutes(handler handlers.FcmNotificationHandlerInterface) *FcmNotificationRoutes {
	return &FcmNotificationRoutes{
		handler: handler,
	}
}

func (h *FcmNotificationRoutes) Setup(r *gin.Engine) {
	route := r.Group("notification")
	{
		route.GET("create/token", middlewares.CheckAccessToken(), h.handler.Create)
		route.POST("noti", middlewares.CheckAccessToken(), h.handler.SentNoti)
	}
}

func (h *FcmNotificationRoutes) CreatePushNotificationCronjob() {
	h.handler.CreateCronjob()
}

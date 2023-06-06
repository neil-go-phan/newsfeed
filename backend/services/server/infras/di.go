package infras

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func SetupRoute(db *gorm.DB, r *gin.Engine) {
	userRoutes := InitizeUser(db)

	userRoutes.Setup(r)
}
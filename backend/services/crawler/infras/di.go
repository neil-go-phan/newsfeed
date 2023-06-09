package infras

import (
	"gorm.io/gorm"
)
func Listen(db *gorm.DB, port string) {
	gRPCServer := InitizeGRPCServer(db)
	gRPCServer.Listen(port)
}
// go:build wireinject
// + build wireinject
// go:generate go run github.com/google/wire/cmd/wire@latest
package infras

import (
	"server/handlers"
	"server/repository"
	"server/routes"
	"server/services"
	"server/services/user"
	"server/services/role"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitUser(db *gorm.DB) *routes.UserRoutes {
	wire.Build(
		repository.NewUserRepo,
		repository.NewRoleRepo,
		userservice.NewUserService,
		roleservice.NewRoleService,
		handlers.NewUserHandler,
		routes.NewUserRoutes,
		wire.Bind(new(repository.UserRepository), new(*repository.UserRepo)),
		wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepo)),
		wire.Bind(new(services.UserServices), new(*userservice.UserService)),
		wire.Bind(new(services.RoleServices), new(*roleservice.RoleService)),
		wire.Bind(new(handlers.UserHandlerInterface), new(*handlers.UserHandler)),
	)
	return &routes.UserRoutes{}
}
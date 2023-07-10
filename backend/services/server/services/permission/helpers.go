package permissionservice

import (
	"server/entities"
	"server/services"
)

func castPermissionToResponse(permission entities.Permission) services.PermissionResponse {
	return services.PermissionResponse{
		ID: permission.ID,
		Entity: permission.Entity,
		Method: permission.Method,
		Description: permission.Description,
	}
}
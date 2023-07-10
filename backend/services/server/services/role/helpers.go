package roleservice

import (
	"server/entities"
	"server/services"

	"gorm.io/gorm"
)

func castRolesToResponse(role entities.Role) services.RoleResponse {
	permissionResponse := make([]services.PermissionResponse, 0)
	for _, permission := range role.Permissions {
		newPermission := services.PermissionResponse{
			ID: permission.ID,
			Entity: permission.Entity,
			Method: permission.Method,
			Description: permission.Description,
		}
		permissionResponse = append(permissionResponse, newPermission)
	}

	return services.RoleResponse{
		ID: role.ID,
		Name: role.Name,
		Description: role.Description,
		Permissions: permissionResponse,
	}
}

func castRoleCreatePayloadToEntity(rolePayload services.RoleResponse ) entities.Role {
	permissions := make([]*entities.Permission, 0)
	for _, permission := range rolePayload.Permissions {
		newPermission := entities.Permission{
			Model: gorm.Model{
				ID: permission.ID,
			},
			Entity: permission.Entity,
			Method: permission.Method,
			Description: permission.Description,
		}
		permissions = append(permissions, &newPermission)
	}

	return entities.Role{
		Name: rolePayload.Name,
		Description: rolePayload.Description,
		Permissions: permissions,
	}
}
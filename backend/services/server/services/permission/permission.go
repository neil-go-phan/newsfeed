package permissionservice

import (
	"server/repository"
	"server/services"
)

type PermissionService struct {
	repo repository.PermissionRepository
}

func NewPermissionService(r repository.PermissionRepository) *PermissionService {
	return &PermissionService{
		repo: r,
	}
}

func (s *PermissionService) List() ([]services.PermissionResponse, error) {
	permissionResponse := make([]services.PermissionResponse, 0)
	permissions, err := s.repo.List()
	if err != nil {
		return permissionResponse, err
	}
	for _, permission := range permissions {
		permissionResponse = append(permissionResponse, castPermissionToResponse(permission))
	}
	return permissionResponse, nil 
}


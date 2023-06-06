package roleservice

import (
	"backend/entities"
	"backend/repository"
	"errors"
)

type RoleService struct {
	repo repository.RoleRepository
}


func NewRoleService(r repository.RoleRepository) *RoleService {
	return &RoleService{
		repo: r,
	}
}

func (s *RoleService) GetRole(roleName string) (*entities.Role, error) {
	return s.repo.Get(roleName)
}

func (s *RoleService) ListRole() (role *[]entities.Role, err error) {
	return s.repo.List()
}

func (s *RoleService) Validate(roleName string) (err error) {
	role, err := s.repo.Get(roleName)
	if err != nil {
		return err
	}
	if role.Name == "" {
		return errors.New("role invalid")
	}
	return nil
}
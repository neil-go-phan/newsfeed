package roleservice

import (
	"errors"
	"fmt"
	"server/repository"
	"server/services"

	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
)

type RoleService struct {
	repo repository.RoleRepository
}

const ROLE_ROLE_ENTITY = "ROLE"
const ROLE_ROLE_CREATE_METHOD = "CREATE"
const ROLE_ROLE_UPDATE_METHOD = "UPDATE"
const ROLE_ROLE_DELETE_METHOD = "DELETE"

func NewRoleService(r repository.RoleRepository) *RoleService {
	return &RoleService{
		repo: r,
	}
}

func (s *RoleService) Create(rolePayload services.RoleResponse) error {
	validate := validator.New()
	err := validate.Struct(rolePayload)
	if err != nil {
		return err
	}
	role := castRoleCreatePayloadToEntity(rolePayload)
	err = s.repo.Create(role)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoleService) Get(roleName string) (services.RoleResponse, error) {
	roleResponse := services.RoleResponse{}
	role, err := s.repo.Get(roleName)
	if err != nil {
		return roleResponse, err
	}

	roleResponse = castRolesToResponse(role)
	return roleResponse, nil
}

func (s *RoleService) List(page int, pageSize int) ([]services.RoleResponse, error) {
	roleResponse := make([]services.RoleResponse, 0)
	roles, err := s.repo.List(page, pageSize)
	if err != nil {
		return roleResponse, err
	}
	for _, role := range roles {
		roleResponse = append(roleResponse, castRolesToResponse(role))
	}
	return roleResponse, nil
}

func (s *RoleService) Count() (int, error) {
	return s.repo.Count()
}

func (s *RoleService) ListRoleName() ([]string, error)  {
	names := make([]string, 0)
	roles, err := s.repo.ListRoleName()
	if err != nil {
		return names, err
	}
	for _, role := range roles {
		names = append(names, role.Name)
	}
	return names, nil
}

func (s *RoleService) Delete(id uint) (error) {
	if (id > 0 && id < 4) {
		return fmt.Errorf("cant not delete default roles")
	}

	return s.repo.Delete(id)
}

func (s *RoleService) Update(rolePayload services.RoleResponse) error {
	validate := validator.New()
	err := validate.Struct(rolePayload)
	if err != nil {
		return err
	}
	role := castRoleCreatePayloadToEntity(rolePayload)
	role.ID = rolePayload.ID
	err = s.repo.Update(role)
	if err != nil {
		return err
	}
	return nil
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

func (s *RoleService) GrantPermission(userRole string, entity string, method string) bool {
	role, err := s.repo.Get(userRole)
	if err != nil {
		log.Error(err)
		return false
	}
	for _, permission := range role.Permissions {
		if permission.Entity == entity && permission.Method == method {
			return true
		}
	}
	return false
}

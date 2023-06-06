package repository

import (
	"backend/entities"

	"gorm.io/gorm"
)

type RoleRepo struct {
	DB *gorm.DB
}

type RoleRepository interface {
	Get(roleName string) (r *entities.Role, err error)
	List() (role *[]entities.Role, err error)
}

func NewRoleRepo(db *gorm.DB) *RoleRepo {
	return &RoleRepo{
		DB: db,
	}
}

func (repo *RoleRepo) Get(roleName string) (r *entities.Role, err error) {
	role := new(entities.Role)

	err = repo.DB.Where("role_name = ?", roleName).Find(&role).Error
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (repo *RoleRepo) List() (role *[]entities.Role, err error) {
	roles := make([]entities.Role, 10)

	err = repo.DB.Select("role_name, description").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	
	return &roles, nil
}

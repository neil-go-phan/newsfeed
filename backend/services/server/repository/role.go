package repository

import (
	"server/entities"
	"server/helpers"

	"gorm.io/gorm"
)

type RoleRepo struct {
	DB *gorm.DB
}

//go:generate mockery --name RoleRepository
type RoleRepository interface {
	Get(roleName string) (entities.Role, error)
	List(page int, pageSize int) ([]entities.Role, error)
	Create(role entities.Role) error
	Update(role entities.Role) error
	ListRoleName() ([]entities.Role, error) 
	Delete(id uint) error
	Count() (int, error)
}


func NewRoleRepo(db *gorm.DB) *RoleRepo {
	return &RoleRepo{
		DB: db,
	}
}

func (repo *RoleRepo) Create(role entities.Role) error {
	err := repo.DB.Create(&role).Save(&role).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *RoleRepo) Get(roleName string) (entities.Role, error) {
	role := new(entities.Role)

	err := repo.DB.Preload("Permissions").Where("name = ?", roleName).Find(&role).Error
	if err != nil {
		return *role, err
	}

	return *role, nil
}

func (repo *RoleRepo) List(page int, pageSize int) ([]entities.Role, error) {
	roles := make([]entities.Role, 0)

	err := repo.DB.Preload("Permissions").Scopes(helpers.Paginate(page, pageSize)).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (repo *RoleRepo) ListRoleName() ([]entities.Role, error) {
	roles := make([]entities.Role, 0)

	err := repo.DB.Select("name").Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (repo *RoleRepo) Count() (int, error) {
	var count int64
	err := repo.DB.Table("roles").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (repo *RoleRepo) Delete(id uint) error {
	err := repo.DB.Unscoped().Delete(&entities.Role{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *RoleRepo) Update(role entities.Role) error {
	err := repo.DB.Model(&role).
		Where("id = ?", role.ID).
		Updates(entities.Role{
			Name:        role.Name,
			Description: role.Description,
		}).Error
	if err != nil {
		return err
	}
	err = repo.DB.Model(&role).
		Association("Permissions").
		Replace(role.Permissions)
	if err != nil {
		return err
	}

	return nil
}

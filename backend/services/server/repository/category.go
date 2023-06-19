package repository

import (
	"fmt"
	"server/entities"
	"server/helpers"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Get(name string) (entities.Category, error)
	List() ([]entities.Category, error)
	GetPagination(page int, pageSize int) ([]entities.Category, error)
	Count() (int, error)

	CreateIfNotExist(category entities.Category) error
	Delete(category entities.Category) error
	UpdateName(category entities.Category, newName string) error
}

type CategoryRepo struct {
	DB *gorm.DB
}

func NewCategory(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{
		DB: db,
	}
}

func (repo *CategoryRepo) CreateIfNotExist(category entities.Category) error {
	result := repo.DB.Where(entities.Category{Name: category.Name}).FirstOrCreate(&category)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("category already exist")
	}

	return nil
}

func (repo *CategoryRepo) Get(name string) (entities.Category, error) {
	category := new(entities.Category)
	err := repo.DB.Where("name = ? ", name).First(category).Error
	if err != nil {
		return *category, err
	}
	return *category, nil
}

func (repo *CategoryRepo) List() ([]entities.Category, error) {
	categories := make([]entities.Category, 0)
	err := repo.DB.Find(&categories).Error
	if err != nil {
		return categories, err
	}
	return categories, nil
}

func (repo *CategoryRepo) Delete(category entities.Category) error {
	err := repo.DB.
		Where("name = ?", category.Name).
		Unscoped().
		Delete(&category).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *CategoryRepo) UpdateName(category entities.Category, newName string) error {
	err := repo.DB.Model(&category).
		Update("name", newName).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *CategoryRepo) GetPagination(page int, pageSize int) ([]entities.Category, error) {
	categories := make([]entities.Category, 0)

	err := repo.DB.Scopes(helpers.Paginate(page, pageSize)).Find(&categories).Error
	if err != nil {
		return categories, err
	}
	return categories, nil
}

func (repo *CategoryRepo) Count() (int, error) {
	var count int64
	err := repo.DB.Table("categories").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

package repository

import (
	"fmt"
	"server/entities"
	"server/helpers"

	"gorm.io/gorm"
)

type TopicRepository interface {
	List() ([]entities.Topic, error)
	GetPagination(page int, pageSize int) ([]entities.Topic, error)
	Count() (int, error)
	
	CreateIfNotExist(topic entities.Topic) error
	Delete(topic entities.Topic) error
	Update(topic entities.Topic) error
	UpdateWhenDeteleCategory(oldCategoryID uint, newCategoryID uint) error
}

type TopicRepo struct {
	DB *gorm.DB
}

func NewTopic(db *gorm.DB) *TopicRepo {
	return &TopicRepo{
		DB: db,
	}
}

func (repo *TopicRepo) CreateIfNotExist(topic entities.Topic) error {
	result := repo.DB.Where(entities.Topic{Name: topic.Name}).FirstOrCreate(&topic)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("topic already exist")
	}

	return nil
}

func (repo *TopicRepo) List() ([]entities.Topic, error) {
	topics := make([]entities.Topic, 0)
	err := repo.DB.Find(&topics).Error
	if err != nil {
		return topics, err
	}
	return topics, nil
}

func (repo *TopicRepo) Delete(topic entities.Topic) error {
	err := repo.DB.
		Where("name = ?", topic.Name).
		Unscoped().
		Delete(&topic).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *TopicRepo) Update(topic entities.Topic) error {
	err := repo.DB.Model(&topic).Updates(entities.Topic{Name: topic.Name, CategoryID: topic.CategoryID}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *TopicRepo) UpdateWhenDeteleCategory(oldCategoryID uint, newCategoryID uint) error {
	err := repo.DB.Model(&entities.Topic{}).
		Where("category_id = ?", oldCategoryID).
		Update("category_id", newCategoryID).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *TopicRepo) GetPagination(page int, pageSize int) ([]entities.Topic, error) {
	topics := make([]entities.Topic, 0)

	err := repo.DB.Scopes(helpers.Paginate(page, pageSize)).Find(&topics).Error
	if err != nil {
		return topics, err
	}
	return topics, nil
}

func (repo *TopicRepo) Count() (int, error) {
	var count int64
	err := repo.DB.Table("topics").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

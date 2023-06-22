package categoryservices

import (
	"fmt"
	"server/entities"
	"server/repository"
	"server/services"
	"strings"
)

const OTHERS_CATEGORY_ID = 1
const OTHERS_CATEGORY_NAME = "Others"
const ERROR_RAISE_WHEN_DELETE_A_CATEGORY_THEN_CREATE_IT_AGAIN = `duplicate key value violates unique constraint "categories_name_key"`

type CategoryService struct {
	repo          repository.CategoryRepository
	topicServices services.TopicServices
}

func NewCategoryService(repo repository.CategoryRepository, topicServices services.TopicServices) *CategoryService {
	categoryService := &CategoryService{
		repo:          repo,
		topicServices: topicServices,
	}
	return categoryService
}

func (s *CategoryService) CreateIfNotExist(category entities.Category) error {
	category.Name = strings.TrimSpace(category.Name)
	err := validateCategoryName(category)
	if err != nil {
		return err
	}

	err = s.repo.CreateIfNotExist(category)
	if err != nil {
		return err
	}
	return nil
}

func (s *CategoryService) ListName() ([]services.CategoryResponse, error) {
	categoriesResponse := make([]services.CategoryResponse, 0)
	categories, err := s.repo.ListName()
	if err != nil {
		return categoriesResponse, err
	}
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, castEntityCategoryToCategoryResponse(category))
	}
	return categoriesResponse, nil
}

func (s *CategoryService) ListAll() ([]services.CategoryResponse, error) {
	categoriesResponse := make([]services.CategoryResponse, 0)
	categories, err := s.repo.ListAll()
	if err != nil {
		return categoriesResponse, err
	}
	for _, category := range categories {
		if !isOthersCategory(category) {
			categoriesResponse = append(categoriesResponse, castEntityCategoryToCategoryResponse(category))
		}
	}
	return categoriesResponse, nil
}

func (s *CategoryService) UpdateName(payload services.UpdateNameCategoryPayload) error {
	category, newName := extractUpdateNamePayload(payload)
	err := validateCategoryName(category)
	if err != nil {
		return err
	}
	checkCategory, err := s.repo.Get(newName)
	if checkCategory.Name == newName && err == nil {
		return fmt.Errorf("category %s already exist", newName)
	}
	return s.repo.Update(category, newName)
}

func (s *CategoryService) Delete(category entities.Category) error {
	category.Name = strings.TrimSpace(category.Name)
	err := validateCategoryName(category)
	if err != nil {
		return err
	}
	if category.ID == 0 {
		return fmt.Errorf("not found category id")
	}
	err = s.topicServices.UpdateWhenDeteleCategory(category.ID, OTHERS_CATEGORY_ID)
	if err != nil {
		return err
	}
	return s.repo.Delete(category)
}

func (s *CategoryService) GetPagination(page int, pageSize int) ([]services.CategoryResponse, error) {
	categoriesResponse := make([]services.CategoryResponse, 0)
	categories, err := s.repo.GetPagination(page, pageSize)
	if err != nil {
		return categoriesResponse, err
	}
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, castEntityCategoryToCategoryResponse(category))
	}
	return categoriesResponse, nil
}

func (s *CategoryService) Count() (int, error) {
	return s.repo.Count()
}
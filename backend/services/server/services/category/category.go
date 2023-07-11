package categoryservices

import (
	"fmt"
	"server/entities"
	"server/repository"
	"server/services"
	"strings"

	"gorm.io/gorm"
)

const OTHERS_CATEGORY_ID = 1
const OTHERS_CATEGORY_NAME = "Others"
const ERROR_RAISE_WHEN_DELETE_A_CATEGORY_THEN_CREATE_IT_AGAIN = `duplicate key value violates unique constraint "categories_name_key"`
const CATEGORY_ROLE_ENTITY = "CATEGORY"
const CATEGORY_ROLE_CREATE_METHOD = "CREATE"
const CATEGORY_ROLE_UPDATE_METHOD = "UPDATE"
const CATEGORY_ROLE_DELETE_METHOD = "DELETE"

type CategoryService struct {
	repo          repository.CategoryRepository
	topicServices services.TopicServices
	roleServices  services.RoleServices
}

func NewCategoryService(repo repository.CategoryRepository, topicServices services.TopicServices, roleServices services.RoleServices) *CategoryService {
	categoryService := &CategoryService{
		repo:          repo,
		topicServices: topicServices,
		roleServices:  roleServices,
	}
	return categoryService
}

func (s *CategoryService) CreateIfNotExist(role string, category entities.Category) error {
	isAllowed := s.roleServices.GrantPermission(role, CATEGORY_ROLE_ENTITY, CATEGORY_ROLE_CREATE_METHOD)
	if !isAllowed {
		return fmt.Errorf("unauthorized")
	}
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

func (s *CategoryService) Update(role string, payload services.UpdateNameCategoryPayload) error {
	isAllowed := s.roleServices.GrantPermission(role, CATEGORY_ROLE_ENTITY, CATEGORY_ROLE_UPDATE_METHOD)
	if !isAllowed {
		return fmt.Errorf("unauthorized")
	}

	categoty := entities.Category{
		Model: gorm.Model{
			ID: payload.Category.ID,
		},
		Name: payload.NewName,
		Illustration: payload.NewIllustration,
	}
	
	return s.repo.Update(categoty)
}

func (s *CategoryService) Delete(role string, category entities.Category) error {
	isAllowed := s.roleServices.GrantPermission(role, CATEGORY_ROLE_ENTITY, CATEGORY_ROLE_DELETE_METHOD)
	if !isAllowed {
		return fmt.Errorf("unauthorized")
	}
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

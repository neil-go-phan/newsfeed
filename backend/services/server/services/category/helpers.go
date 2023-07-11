package categoryservices

import (
	"fmt"
	"server/entities"
	"server/services"

	"github.com/go-playground/validator"
)

func validateCategoryName(category entities.Category) error {
	if isOthersCategory(category) {
		return fmt.Errorf("can not change 'Others' category")
	}
	validate := validator.New()
	err := validate.Struct(category)
	if err != nil {
		return err
	}

	return nil
}

func isOthersCategory(category entities.Category) bool {
	return category.Name == OTHERS_CATEGORY_NAME || category.ID == OTHERS_CATEGORY_ID 
} 

func castEntityCategoryToCategoryResponse(entityCategory entities.Category) services.CategoryResponse {
	return services.CategoryResponse{
		ID:           entityCategory.ID,
		Name:         entityCategory.Name,
		Illustration: entityCategory.Illustration,
	}
}

package categoryservices

import (
	"fmt"
	"server/entities"
	"server/services"
	"strings"

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

func extractUpdateNamePayload(payload services.UpdateNameCategoryPayload) (entities.Category, string) {
	category := entities.Category{
		Name:      strings.TrimSpace(payload.Category.Name),
		Illustration: payload.Category.Illustration,
	}
	newName := payload.NewName
	category.ID = payload.Category.ID
	return category, newName
}

func castEntityCategoryToCategoryResponse(entityCategory entities.Category) services.CategoryResponse {
	return services.CategoryResponse{
		ID:           entityCategory.ID,
		Name:         entityCategory.Name,
		Illustration: entityCategory.Illustration,
	}
}

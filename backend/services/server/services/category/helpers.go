package categoryservices

import (
	"server/entities"
	"server/services"
	"strings"

	"github.com/go-playground/validator"
)

func validateCategory(category entities.Category) (error) {
	validate := validator.New()
	err := validate.Struct(category)
	if err != nil {
		return err
	}

	return nil
}

func extractUpdateNamePayload(payload services.UpdateNameCategoryPayload) (entities.Category, string) {
	category := entities.Category{
		Name: strings.TrimSpace(payload.Category.Name),
	}
	newName := payload.NewName
	category.ID = payload.Category.ID
	return category, newName
}

func castEntityCategoryToCategoryResponse(entityCategory entities.Category) (services.CategoryResponse) {
	return services.CategoryResponse{
		Name: entityCategory.Name,
		ID: entityCategory.ID,
	}
}
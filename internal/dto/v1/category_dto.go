package v1dto

import "github.com/dangLuan01/rebuild-api-movie28/internal/models"

type CategoryDTO struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func MapCategoryDTO(categories []models.Category) []CategoryDTO {

	category_dto := make([]CategoryDTO, 0, len(categories))
	for _, category:= range categories {
		cate := CategoryDTO {
			Name: category.Name,
			Slug: category.Slug,
		}
		category_dto = append(category_dto, cate)
	}

	return category_dto
}
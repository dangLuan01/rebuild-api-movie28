package v1dto

import "github.com/dangLuan01/rebuild-api-movie28/internal/models"

type GenreDTO struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Image    string `json:"image"`
}

func MapGenreDTO(genre models.Genre) *GenreDTO {
	return &GenreDTO{
		Name: genre.Name,
		Slug: genre.Slug,
		Image: genre.Image,
	}
}

func MapGenresDTO(genres []models.Genre) []GenreDTO {
	dtos := make([]GenreDTO, 0, len(genres))
	for _, genre := range genres {
		dtos = append(dtos, *MapGenreDTO(genre))
	}

	return dtos
}
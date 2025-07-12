package v1dto

import "github.com/dangLuan01/rebuild-api-movie28/internal/models"

type GenreDTO struct {
	Name     string `json:"name"`
	Slug     string `json:"slug,omitempty"`
	Image    string `json:"image,omitempty"`
	Total 	 int 	`json:"total,omitempty"`
}
type GenreWithMovieDTO struct {
	Genre 		GenreDTO `json:"genre"`
	Movies 		[]MovieDTO `json:"movies"`
	Page 		int64 `json:"page"`
	PageSize 	int64 `json:"page_size"`
	TotalPages 	int64 `json:"total_pages"`
}
func MapGenreWithMovie(genre models.GenreWithMovie) *GenreWithMovieDTO {
	movie_dto := make([]MovieDTO, 0, len(genre.Movie))
	for _, movie := range genre.Movie {
		m := MovieDTO{
			Name: movie.Name,
			Origin_name: movie.Origin_name,
			Slug: movie.Slug,
			Type: movie.Type,
			Release_date: movie.Release_date,
			Rating: movie.Rating,
			Image: ImageDTO{
				Poster: movie.Image.Poster,
			},
			Genres: []GenreDTO{
				{
					Name: genre.Genre.Name,
				},
			},
		}
		movie_dto = append(movie_dto, m)
	}
	return &GenreWithMovieDTO{
		Genre: GenreDTO{
			Name: genre.Genre.Name,
		},
		Movies: movie_dto,
		Page: genre.Page,
		PageSize: genre.PageSize,
		TotalPages: genre.TotalPages,
	}
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
func MapGenresWithTotalDTO(genres []models.Genre) []GenreDTO {
	dtos := make([]GenreDTO, 0, len(genres))
	for _, genre := range genres {
		g := GenreDTO{
			Name: genre.Name,
			Slug: genre.Slug,
			Image: genre.Image,
			Total: genre.Total,
		}
		dtos = append(dtos, g)
	}

	return dtos
}
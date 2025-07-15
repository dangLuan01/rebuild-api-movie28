package v1dto

import (
	"github.com/dangLuan01/rebuild-api-movie28/internal/models"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
)

type GenreDTO struct {
	Name     string `json:"name"`
	Slug     string `json:"slug,omitempty"`
	Image    string `json:"image,omitempty"`
	Total 	 int 	`json:"total,omitempty"`
}
type GenreWithMovieDTO struct {
	Genre 		GenreDTO 		`json:"genre"`
	Movies 		[]MovieDTO 		`json:"movies"`
	Paginate 	Paginate 		`json:"paginate"`
}
func MapGenreWithMovie(movies []MovieRawDTO, genre models.Genre, paginate Paginate) *GenreWithMovieDTO {
	movie_dto := make([]MovieDTO, 0, len(movies))
	for _, movie := range movies {
		m := MovieDTO{
			Name: movie.Name,
			Origin_name: movie.Origin_name,
			Slug: movie.Slug,
			Type: movie.Type,
			Release_date: movie.Release_date,
			Rating: utils.ConvertRating(float32(movie.Rating)),
			Image: ImageDTO{
				Poster: movie.Poster,
			},
			Genres: []GenreDTO{
				{
					Name: genre.Name,
				},
			},
			Episode: movie.Episode,
			Episode_total: movie.Episode_total,
		}
		movie_dto = append(movie_dto, m)
	}
	return &GenreWithMovieDTO{
		Genre: GenreDTO{
			Name: genre.Name,
		},
		Movies: movie_dto,
		Paginate: Paginate{
			Page: paginate.Page,
			PageSize: paginate.PageSize,
			TotalPages: paginate.TotalPages,
		},
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
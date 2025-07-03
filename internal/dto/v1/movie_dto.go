package v1dto

import (
	"github.com/dangLuan01/rebuild-api-movie28/internal/models"
)

type MovieDTO struct {
	Name         string     `json:"name"`
	Origin_name  string     `json:"origin_name"`
	Slug         string     `json:"slug"`
	Type         string     `json:"type"`
	Release_date int        `json:"release_date"`
	Rating       float32    `json:"rating"`
	Content      string     `json:"content,omitempty"`
	Runtime      string     `json:"runtime,omitempty"`
	Age          string     `json:"age,omitempty"`
	Trailer      string     `json:"trailer,omitempty"`
	Image        ImageDTO   `json:"image"`
	Genres       []GenreDTO `json:"genres"`
}
type MovieRawDTO struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Origin_name  string  `json:"origin_name"`
	Slug         string  `json:"slug"`
	Type         string  `json:"type"`
	Release_date int     `json:"release_date"`
	Rating       float64 `json:"rating"`
	Content      string  `json:"content,omitempty"`
	Runtime      string  `json:"runtime,omitempty"`
	Age          string  `json:"age,omitempty"`
	Trailer      string  `json:"trailer,omitempty"`
	Thumb        string  `json:"thumb"`
	Poster       string  `json:"poster"`
	Genre        string  `json:"genre"`
}
type ImageDTO struct {
	Thumb  string `json:"thumb,omitempty"`
	Poster string `json:"poster,omitempty"`
}

func MapMovieDTO(movies []models.Movie) []MovieDTO {
	movie_dto := make([]MovieDTO, 0, len(movies))
	for _, movie := range movies {
		m := MovieDTO{
			Name: movie.Name,
			Origin_name: movie.Origin_name,
			Slug: movie.Slug,
			Type: movie.Type,
			Release_date: movie.Release_date,
			Rating: movie.Rating,
			Image: ImageDTO{
				Thumb: movie.Image.Thumb,
			},
			Genres: []GenreDTO{
				{
					Name: movie.Genres[0].Name,
				},
			},	
		}
		
		movie_dto = append(movie_dto, m)
	}

	return movie_dto
}

func MapMovieRawToMovie(movieRaw []MovieRawDTO) []models.Movie {
	movie_dto := make([]models.Movie, 0, len(movieRaw))
	for _, movie := range movieRaw {
		movie := models.Movie {
			Name: movie.Name,
			Origin_name: movie.Origin_name,
			Slug: movie.Slug,
			Type: movie.Type,
			Release_date: movie.Release_date,
			Rating: float32(movie.Rating),
			Image: models.Image {
				Thumb: movie.Thumb,
			},
			Genres: []models.Genre {
				{
					Name: movie.Genre,
				},
			},
			
		}
		movie_dto = append(movie_dto, movie)
	}
	return movie_dto
}
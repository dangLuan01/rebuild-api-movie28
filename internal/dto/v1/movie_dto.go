package v1dto

import (
	"time"

	"github.com/dangLuan01/rebuild-api-movie28/internal/models"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
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
	Genres       []GenreDTO `json:"genres,omitempty"`
	Episode		 int 		`json:"episode"`
	Episode_total 		 int 		`json:"episode_total"`
}
type MovieRawDTO struct {
	Id           int     	`json:"id"`
	Name         string  	`json:"name"`
	Origin_name  string  	`json:"origin_name"`
	Slug         string  	`json:"slug"`
	Type         string  	`json:"type"`
	Release_date int     	`json:"release_date"`
	Rating       float64 	`json:"rating"`
	Content      string  	`json:"content,omitempty"`
	Runtime      string  	`json:"runtime,omitempty"`
	Age          string  	`json:"age,omitempty"`
	Trailer      string  	`json:"trailer,omitempty"`
	Thumb        string  	`json:"thumb"`
	Poster       string  	`json:"poster"`
	Updated_at 	 time.Time 	`json:"updated_at"`
	Genre        string  	`json:"genre"`
	Episode 	 int 	 	`json:"episode"`
	Episode_total 		 int 	 `json:"episode_total"`
}

type ImageDTO struct {
	Thumb  string `json:"thumb,omitempty"`
	Poster string `json:"poster,omitempty"`
}

type EpisodeDTO struct {
	Server_id int    `json:"server_id"`
	Episode   string `json:"episode"`
	Hls       string `json:"hls"`
}

type ServerDTO struct {
	Id       int          `json:"id"`
	Name     string       `json:"name"`
	Episodes []EpisodeDTO `json:"episodes"`
}

type MoviesDTOWithPaginate struct {
	Movie    []MovieDTO `json:"movies"`
	Paginate Paginate   `json:"paginate"`
}

type Paginate struct {
	Page       int64   `json:"page"`
	PageSize   int64   `json:"page_size"`
	TotalPages int64 `json:"total_pages"`
}

type MovieDetailDTO struct {
	Name         string      `json:"name"`
	Origin_name  string      `json:"origin_name"`
	Slug         string      `json:"slug"`
	Type         string      `json:"type"`
	Release_date int         `json:"release_date"`
	Rating       float32     `json:"rating"`
	Content      string      `json:"content,omitempty"`
	Runtime      string      `json:"runtime,omitempty"`
	Age          string      `json:"age,omitempty"`
	Trailer      string      `json:"trailer,omitempty"`
	Updated_at   time.Time	 `json:"updated_at"`
	Image        ImageDTO    `json:"image"`
	Genres       []GenreDTO  `json:"genres"`
	Servers      []ServerDTO `json:"servers"`
}
type Filter struct {
	Genre 			*string
	Release_date 	*string
	Type 			*string
}
func MapMovieDetailDTO(movie MovieRawDTO) *MovieDetailDTO {
	return &MovieDetailDTO{
		Name:         movie.Name,
		Origin_name:  movie.Origin_name,
		Slug:         movie.Slug,
		Type:         movie.Type,
		Release_date: movie.Release_date,
		Rating:       utils.ConvertRating(float32(movie.Rating)),
		Content:      movie.Content,
		Runtime:      movie.Runtime,
		Age:          movie.Age,
		Trailer:      movie.Trailer,
		Updated_at:   movie.Updated_at,
		Image: ImageDTO{
			Thumb:  movie.Thumb,
			Poster: movie.Poster,
		},
	}
}
func MapMovieDTO(movies MovieRawDTO) *MovieDTO {
	return &MovieDTO{
		Name:         movies.Name,
		Origin_name:  movies.Origin_name,
		Slug:         movies.Slug,
		Type:         movies.Type,
		Release_date: movies.Release_date,
		Rating:       utils.ConvertRating(float32(movies.Rating)),
		Episode_total: 		  movies.Episode_total,
		Episode: 	  movies.Episode,	
		Image: ImageDTO{
			Thumb:  movies.Thumb,
			Poster: movies.Poster,
		},
		Genres: []GenreDTO{
			{
				Name: movies.Genre,
			},
		},
		
	}
}

func MapMovieRawToMovieDTO(movieRaw []MovieRawDTO) []MovieDTO {
	movie_dto := make([]MovieDTO, 0, len(movieRaw))
	for _, movie := range movieRaw {

		movie_dto = append(movie_dto, *MapMovieDTO(movie))
	}

	return movie_dto
}

func MapMovieDTOWithPanigate(movies []MovieRawDTO, paginate Paginate) *MoviesDTOWithPaginate {
	return &MoviesDTOWithPaginate{
		Movie: MapMovieRawToMovieDTO(movies),
		Paginate: Paginate{
			paginate.Page,
			paginate.PageSize,
			paginate.TotalPages,
		},
	}
}

func MapMovieModelTODTO(movies []models.Movie) []MovieDTO {
	movie_dto := make([]MovieDTO, 0, len(movies))
	for _, movie := range movies {
		m := MovieDTO{
			Name: movie.Name,
			Origin_name: movie.Origin_name,
			Slug: movie.Slug,
			Image: ImageDTO{
				Poster: movie.Image.Poster,
			},
			Type: movie.Type,
			Age: movie.Age,
			Release_date: movie.Release_date,
			Runtime: movie.Runtime,
		}

		movie_dto = append(movie_dto, m)
	}
	
	return movie_dto
}
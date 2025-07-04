package v1dto

type MovieDTO struct {
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
	Image        ImageDTO    `json:"image"`
	Genres       []GenreDTO  `json:"genres"`
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
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
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
	Image        ImageDTO    `json:"image"`
	Genres       []GenreDTO  `json:"genres"`
	Servers      []ServerDTO `json:"servers"`
}

func MapMovieDetailDTO(movie MovieRawDTO) *MovieDetailDTO {
	return &MovieDetailDTO{
		Name:         movie.Name,
		Origin_name:  movie.Origin_name,
		Slug:         movie.Slug,
		Type:         movie.Type,
		Release_date: movie.Release_date,
		Rating:       float32(movie.Rating),
		Content:      movie.Content,
		Runtime:      movie.Runtime,
		Age:          movie.Age,
		Trailer:      movie.Trailer,
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
		Rating:       float32(movies.Rating),
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
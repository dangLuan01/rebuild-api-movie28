package v1dto

type ThemeDTO struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Genre_id   *int    `json:"genre_id,omitempty"`
	Country_id *int    `json:"country_id,omitempty"`
	Type       *string `json:"type,omitempty"`
	Year       *int    `json:"year,omitempty"`
	Limit      int64     `json:"limit"`
	Layout     int     `json:"layout"`
}

type ThemeWithMovieDTO struct {
	Theme  ThemeDTO              `json:"theme"`
	Movies MoviesDTOWithPaginate `json:"movies_of_theme"`
}

type ThemesWithPaginateDTO struct {
	Themes   []ThemeWithMovieDTO `json:"data_themes"`
	Paginate Paginate            `json:"paginate"`
}

func MapThemeDTOWithPaginate(theme []ThemeWithMovieDTO, paginate Paginate) *ThemesWithPaginateDTO {
	return &ThemesWithPaginateDTO{
		Themes:   theme,
		Paginate: paginate,
	}
}
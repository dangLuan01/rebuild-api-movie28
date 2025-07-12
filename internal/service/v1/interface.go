package v1service

import (
	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
	"github.com/dangLuan01/rebuild-api-movie28/internal/models"
)

type UserService interface {
	GetAllUser() ([]models.User, error)
	GetUserByUUID(uuid string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(uuid string, user models.User) (models.User, error)
	DeleteUser(uuid string) error
}

type GenreService interface {
	GetAllGenres() ([]models.Genre, error)
	GetGenreBySlug(slug string, page, page_size int) (models.GenreWithMovie, error)
}

type MovieService interface {
	GetMovieHot(limit int) ([]v1dto.MovieRawDTO, error)
	GetAllMovies(page, pageSize int) ([]v1dto.MovieRawDTO, v1dto.Paginate, error)
	GetMovieDetail(slug string) (*v1dto.MovieDetailDTO, error)
	FilterMovie(filter *v1dto.Filter, page, pageSize int) ([]v1dto.MovieRawDTO, *v1dto.Paginate, error)
}

type CategoryService interface {
	GetAllCategory() ([]models.Category, error)
}

type ThemeService interface {
	GetAllThemes(param ThemeParam) (*v1dto.ThemesWithPaginateDTO, error)
}

type SearchService interface {
	SearchMovie(query string) ([]models.Movie, error)
}
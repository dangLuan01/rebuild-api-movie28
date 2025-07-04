package v1service

import (
	"github.com/dangLuan01/rebuild-api-movie28/internal/models"
	categoryrepository "github.com/dangLuan01/rebuild-api-movie28/internal/repository/category"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
)

type categoryService struct {
	repo categoryrepository.CategoryRepository
}

func NewCategoryService(repo categoryrepository.CategoryRepository) CategoryService {
	return &categoryService{
		repo: repo,
	}
}

func (cs *categoryService) GetAllCategory() ([]models.Category, error) {

	categories, err := cs.repo.FindAll()
	if err != nil {
		return nil, utils.WrapError(
			string(utils.ErrCodeInternal),
			"Faile get all category",
			err,
		)
	}

	return categories, nil
}
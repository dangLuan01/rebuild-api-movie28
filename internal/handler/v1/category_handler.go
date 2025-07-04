package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service v1service.CategoryService
}

func NewCategoryHandler(service v1service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (ch *CategoryHandler) GetAllCategory(ctx *gin.Context) {
	
	categories, err := ch.service.GetAllCategory()
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, v1dto.MapCategoryDTO(categories))
}
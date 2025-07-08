package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
	"github.com/dangLuan01/rebuild-api-movie28/internal/validation"
	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	service v1service.SearchService
}

type GetSearchQuery struct {
	Query    string `form:"query" binding:"min=3"`
}

func NewSearchHandler(service v1service.SearchService) *SearchHandler {
	return &SearchHandler{
		service: service,
	}
}

func (sh *SearchHandler) SeachMovie(ctx *gin.Context) {
	var query GetSearchQuery

	if err := ctx.ShouldBindQuery(&query); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}
	search, err := sh.service.SearchMovie(query.Query)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, v1dto.MapMovieModelTODTO(search))
}
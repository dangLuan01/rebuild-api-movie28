package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
	"github.com/dangLuan01/rebuild-api-movie28/internal/validation"
	"github.com/gin-gonic/gin"
)

type GenreHandler struct {
	service v1service.GenreService
}
type GetGenreBySlugParam struct {
	Slug 		string `uri:"slug" binding:"slug"`
}
type GetGenreBySlugQuery struct {
	Page 		int64 `form:"page" binding:"omitempty,minInt=1"`
	PageSize 	int64 `form:"page_size" binding:"omitempty,minInt=1,maxInt=50"`
}

func NewGenreHandler(service v1service.GenreService) *GenreHandler {
	return &GenreHandler{
		service: service,
	}
}

func (g *GenreHandler) GetAllGenres(ctx *gin.Context) {
	genres, err := g.service.GetAllGenres()
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, v1dto.MapGenresWithTotalDTO(genres))
}
func (g *GenreHandler)GetGenreBySlug(ctx *gin.Context)  {
	var (
		param GetGenreBySlugParam
		query GetGenreBySlugQuery
	)
	errSlug := ctx.ShouldBindUri(&param)
	if errSlug != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(errSlug))
		return
	}
	errQuery := ctx.ShouldBindQuery(&query)
	if errQuery != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(errQuery))
		return
	}
	genre, err := g.service.GetGenreBySlug(param.Slug, query.Page, query.PageSize)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, v1dto.MapGenreWithMovie(genre))
}
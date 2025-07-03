package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
	"github.com/dangLuan01/rebuild-api-movie28/internal/validation"
	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	service v1service.MovieService
}

type GetMovieQuery struct {
	Limit 		int `form:"limit" binding:"omitempty,minInt=1,maxInt=20"`
	Page 		int `form:"page" binding:"omitempty,minInt=1"`
	PageSize 	int `form:"page_size" binding:"omitempty,minInt=1,maxInt=30"`
}

func NewMovieHandler(service v1service.MovieService) *MovieHandler {
	return &MovieHandler{
		service: service,
	}
}

func (mh *MovieHandler) GetMovieWithHot(ctx *gin.Context) {
	var (
		query GetMovieQuery
	)

	errQuery := ctx.ShouldBindQuery(&query)
	if errQuery != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(errQuery))
		return
	}
	movies, err := mh.service.GetMovieHot(query.Limit)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, v1dto.MapMovieRawToMovieDTO(movies))
}

func (mh *MovieHandler) GetAllMovies(ctx *gin.Context)  {
	var (
		query GetMovieQuery
	)
	errQuery := ctx.ShouldBindQuery(&query)
	if errQuery != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(errQuery))
		return
	}
	movies, paginate ,err := mh.service.GetAllMovies(query.Page, query.PageSize)

	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, v1dto.MapMovieDTOWithPanigate(movies, paginate))
}
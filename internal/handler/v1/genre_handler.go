package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
	"github.com/gin-gonic/gin"
)

type GenreHandler struct {
	service v1service.GenreService
}
type GetGenreBySlugParam struct {
	Slug string `uri:"slug" binding:"slug"`
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
	utils.ResponseSuccess(ctx, http.StatusOK, v1dto.MapGenresDTO(genres))
}
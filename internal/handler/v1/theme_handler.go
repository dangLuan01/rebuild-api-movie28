package v1handler

import (
	"net/http"

	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
	"github.com/dangLuan01/rebuild-api-movie28/internal/validation"
	"github.com/gin-gonic/gin"
)

type ThemeHandler struct {
	service v1service.ThemeService
}

type GetThemeQuery struct {
	Id        int `form:"id" binding:"omitempty"`
	Limit     int `form:"limit" binding:"omitempty,minInt=1,maxInt=4"`
	PageTheme int `form:"page_theme" binding:"omitempty,minInt=1"`
	PageMovie int `form:"page_movie" binding:"omitempty,minInt=1,maxInt=20"`
}

func NewThemeHandler(service v1service.ThemeService) *ThemeHandler {
	return &ThemeHandler{
		service: service,
	}
}

func (th *ThemeHandler) GetAllThemes(ctx *gin.Context) {
	var param GetThemeQuery

	if err := ctx.ShouldBindQuery(&param); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

    themes, err := th.service.GetAllThemes(v1service.ThemeParam{
		Id: param.Id,
		PageTheme: param.PageTheme,
		PageMovie: param.PageMovie,
		Limit: param.Limit,
	})
	
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, themes)
	
}
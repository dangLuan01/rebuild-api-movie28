package v1handler

import (
	"strings"

	v1service "github.com/dangLuan01/rebuild-api-movie28/internal/service/v1"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
	"github.com/dangLuan01/rebuild-api-movie28/internal/validation"
	"github.com/gin-gonic/gin"
)

type ProxyHandler struct {
	service v1service.ProxyService
}

type PassHeaderByQuery struct {
	Url string `form:"url" binding:"required"`
}

func NewProxyHandler(service v1service.ProxyService) *ProxyHandler {
	return &ProxyHandler{
		service: service,
	}
}

func (ph *ProxyHandler)PassHeader(ctx *gin.Context) {
	var params PassHeaderByQuery
	if err := ctx.ShouldBindQuery(&params);err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}
	var Url string
	if strings.Contains(params.Url, " ") {
		Url = strings.ReplaceAll(params.Url, " ", "+")
	} else {
		Url = params.Url
	}
	
	ph.service.PassHeader(ctx, Url)
}
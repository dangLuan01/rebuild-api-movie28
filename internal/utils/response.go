package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorCode string

const (
	ErrCodeBadRequest ErrorCode = "BAD_REQUEST"
	ErrCodeNotFound   ErrorCode = "NOT_FOUND"
	ErrCodeConflict   ErrorCode = "CONFLICT"
	ErrCodeInternal   ErrorCode = "INTERNAL_ERROR_SERVER"
)

type AppError struct {
	Code    string
	Message string
	Err     error
}

func (ae *AppError) Error() string {
	return ""
}

func NewError(code, message string) error {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func WrapError(code, message string, err error) error {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func ResponseError(ctx *gin.Context, err error) {
	if appErr, ok := err.(*AppError);ok {
		status := httpStatusFromCode(ErrorCode(appErr.Code))
		response := gin.H{
			"error":	appErr.Message,
			"status":	appErr.Code,
		}

		if appErr.Err != nil {
			response["detail"] = appErr.Err.Error()
		}
		ctx.JSON(status, response)
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
		"status": ErrCodeInternal,
	})
}

func ResponseSuccess(ctx *gin.Context, status int, data any)  {
	ctx.JSON(status, gin.H{
		"status": status,
		"message":"success",
		"data": data,
	})
}

func ResponseSatus(ctx *gin.Context, status int)  {
	ctx.Status(status)
}

func ResponseValidator(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusBadGateway, data)
}
func httpStatusFromCode(code ErrorCode) int {
	switch code {
	case ErrCodeBadRequest:
		return http.StatusBadGateway
	case ErrCodeConflict:
		return http.StatusConflict
	case ErrCodeNotFound:
		return http.StatusNotFound
	default :
		return http.StatusInternalServerError
	}
}

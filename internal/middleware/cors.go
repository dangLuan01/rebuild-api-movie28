package middleware

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context)  {
		
		// origin := ctx.Request.Header.Get("Origin")
		// url := utils.GetEnv("ALLOWED_ORIGIN", "https://www.xoailac.top")
		// if origin == url {
		// 	ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		// 	ctx.Writer.Header().Set("Vary", "Origin")
		// } else {

		// 	ctx.AbortWithStatus(500)
		// 	return
		// }
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, x-api-key")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}
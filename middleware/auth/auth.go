package auth

import (
	"net/http"
	"os"

	"backend_template_personal/common/log"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := os.Getenv("API_KEY")
		if apiKey == "" {
			log.Fatal("missing apikey for authorization")
		}
		apiKeyHeader := ctx.GetHeader("X-API-KEY")
		if apiKeyHeader == "" || apiKeyHeader != apiKey {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorize key is not valid",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

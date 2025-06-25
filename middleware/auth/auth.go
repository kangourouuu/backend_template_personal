package auth

import (
	"net/http"
	"os"

	"backend_template_personal/common/log"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env file can not be load")
	}
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := os.Getenv("API_KEY")
		if apiKey == "" {
			log.Println("missing apikey for authorization")
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

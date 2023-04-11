package middleware

import (
	"clinicaodontologica/pkg/web"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			web.Failure(ctx, http.StatusUnauthorized, errors.New("Token not found"))
			ctx.Abort()
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, http.StatusUnauthorized, errors.New("Invalid token"))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

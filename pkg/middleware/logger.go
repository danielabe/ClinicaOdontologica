package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		verb := ctx.Request.Method
		path := ctx.Request.RequestURI

		ctx.Next()

		statusCode := ctx.Writer.Status()
		fmt.Printf("time: %v\npath: %s\nverb: %s\ncode: %d\n", t, path, verb, statusCode)
	}
}

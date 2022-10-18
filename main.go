package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// config
const port = 8000

func main() {
	// initialize Gin engine
	engine := gin.Default()

	// routing
	engine.GET("/", rootHandler)
	engine.GET("/bye", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Good Bye!")
	})
	engine.GET("/hello.jp", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "こんにちは")
	})
	engine.GET("/hello/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "Hello %s!", name)
	})

	// start server
	engine.Run(fmt.Sprintf(":%d", port))
}

func rootHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello world.")
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*.html")
	engine.GET("/", rootHandler)
	engine.GET("/name-form", nameFormHandler)
	engine.POST("/register-name", registerNameHandler)

	engine.Run(":8000")
}

func rootHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "hello.html", nil)
}

func nameFormHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "name_form.html", nil)
}

func registerNameHandler(ctx *gin.Context) {
	body, _ := ctx.GetRawData()
	ctx.String(http.StatusOK, string(body))
}

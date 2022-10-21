package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}

func NameFormHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "name_form.html", nil)
}

func RegisterNameHandler(ctx *gin.Context) {
	name, _ := ctx.GetPostForm("name")
	ctx.HTML(http.StatusOK, "result.html", gin.H{"Name": name})
}

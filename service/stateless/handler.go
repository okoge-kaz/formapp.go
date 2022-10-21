package stateless

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 安易に小文字で実装すると private になってしまうので注意
func Start(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "start.html", gin.H{"Target": "/stateless/start"})
	// git.H で HTML に対して埋め込みができる
}

func NameForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "name_form.html", gin.H{"Target": "/stateless/name"})
}

func BirthdayForm(ctx *gin.Context) {
	name, exist := ctx.GetPostForm("name")
	if !exist {
		ctx.String(http.StatusBadRequest, "parameter 'name' is not provided")
	}
	ctx.HTML(http.StatusOK, " stateless-birthday-form.html", gin.H{"Name": name, "Target": "/stateless/birthday"})
}

func MessageForm(ctx *gin.Context) {
	name, _ := ctx.GetPostForm("name")
	birthday, _ := ctx.GetPostForm("birthday")

	ctx.HTML(http.StatusOK, "message_form.html", gin.H{
		"Name":     name,
		"Birthday": birthday,
		"Target":   "/stateless/message",
	})
}

func Conformation(ctx *gin.Context) {
	name, _ := ctx.GetPostForm("name")
	birthday, _ := ctx.GetPostForm("birthday")
	message, _ := ctx.GetPostForm("message")

	ctx.HTML(http.StatusOK, "result.html", gin.H{
		"Name":     name,
		"Birthday": birthday,
		"Message":  message,
		"Target":   "/stateless/confirm",
	})
}

func End(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "start.html", gin.H{"Target": "/stateless/start"})
}

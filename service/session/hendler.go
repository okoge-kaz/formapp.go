package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 安易に小文字で実装すると private になってしまうので注意
func Start(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "start.html", gin.H{"Target": "/session/start"})
	// git.H で HTML に対して埋め込みができる
}

func NameForm(ctx *gin.Context) {
	session, err := NewSession()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Fail to create a new session")
		return
	}
	ctx.SetCookie("userID", session.ID(), 600, "/session/", "localhost:8000", false, false)
	ctx.HTML(http.StatusOK, "name_form.html", gin.H{"Target": "/session/name"})
}

func BirthdayForm(ctx *gin.Context) {
	id, err := ctx.Cookie("userID")
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid access")
		return
	}
	session := Session{id}
	name, exist := ctx.GetPostForm("name")
	if !exist {
		ctx.String(http.StatusBadRequest, "parameter 'name' is not provided")
		return
	}
	state, _ := session.GetState()
	state.Name = name
	session.SetState(state)
	ctx.HTML(http.StatusOK, "session-birthday-form.html", nil)
}

func MessageForm(ctx *gin.Context) {
	name, _ := ctx.GetPostForm("name")
	birthday, _ := ctx.GetPostForm("birthday")

	ctx.HTML(http.StatusOK, "message_form.html", gin.H{
		"Name":     name,
		"Birthday": birthday,
		"Target":   "/session/message",
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
		"Target":   "/session/result",
	})
}

func End(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "start.html", gin.H{"Target": "/session/start"})
}

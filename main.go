package main

import (
	"formapp.go/service"
	"formapp.go/service/stateless"
	"formapp.go/service/session"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*.html")

	// root/stateless/*
	engine.GET("/stateless/start", stateless.Start)
	engine.POST("/stateless/start", stateless.NameForm)
	engine.POST("/stateless/name", stateless.BirthdayForm)
	engine.POST("/stateless/birthday", stateless.MessageForm)
	engine.POST("/stateless/message", stateless.Conformation)
	engine.POST("/stateless/result", stateless.End)
	// root/session/*
	engine.GET("/session/start", session.Start)
	engine.POST("/session/start", session.NameForm)
	engine.POST("/session/name", session.BirthdayForm)
	engine.POST("/session/birthday", session.MessageForm)
	engine.POST("/session/message", session.Conformation)
	engine.POST("/session/result", session.End)

	// root/*
	engine.GET("/name-form", service.NameFormHandler)
	engine.POST("/register-name", service.RegisterNameHandler)

	engine.Run(":8000")
}

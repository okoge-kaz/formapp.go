package main

import (
	"formapp.go/service"
	"formapp.go/service/stateless"
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
	// root/*
	engine.GET("/name-form", service.NameFormHandler)
	engine.POST("/register-name", service.RegisterNameHandler)

	engine.Run(":8000")
}

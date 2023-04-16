package router

import (
	"github.com/akiratatsuhisa/first-golang-app/middlewares"
	"github.com/gin-gonic/gin"
)

func Define(r *gin.Engine) {
	r.Use(middlewares.Authenticate())

	authRoute(r.Group(("/auth")))
	todosRoute(r.Group(("/todos")))
}

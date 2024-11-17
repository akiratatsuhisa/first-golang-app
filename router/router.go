package router

import (
	"github.com/akiratatsuhisa/first-golang-app/controllers"
	"github.com/gin-gonic/gin"
)

func Define(r *gin.Engine) {

	r.GET("/sse/:eventName", controllers.SSE)
	r.POST("/send/:eventName", controllers.Send)

	authRoute(r.Group(("/auth")))
	todosRoute(r.Group(("/todos")))
}

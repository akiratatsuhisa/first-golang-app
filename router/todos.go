package router

import (
	"github.com/akiratatsuhisa/first-golang-app/controllers"
	"github.com/akiratatsuhisa/first-golang-app/middlewares"
	"github.com/gin-gonic/gin"
)

func todosRoute(g *gin.RouterGroup) {
	g.Use(middlewares.Authorize())

	g.GET("", middlewares.AuthorizeRoles([]string{"admin"}), controllers.GetTodos)
	g.POST("", controllers.CreateTodo)
	g.GET("/:id", controllers.GetTodo)
	g.PUT("/:id", controllers.UpdateTodo)
	g.DELETE("/:id", controllers.DeleteTodo)
}

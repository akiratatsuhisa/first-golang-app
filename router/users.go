package router

import (
	"github.com/akiratatsuhisa/first-golang-app/controllers"
	"github.com/akiratatsuhisa/first-golang-app/middlewares"
	"github.com/gin-gonic/gin"
)

func usersRoute(g *gin.RouterGroup) {

}

func authRoute(g *gin.RouterGroup) {
	g.POST("/login", controllers.Login)
	g.POST("/register", controllers.Register)
	g.GET("/profile", middlewares.Authorize(), controllers.Profile)
}

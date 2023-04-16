package router

import (
	"github.com/gin-gonic/gin"
)

func Define(r *gin.Engine) {
	todosRoute(r.Group(("/todos")))
}

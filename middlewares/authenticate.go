package middlewares

import (
	"github.com/akiratatsuhisa/first-golang-app/auth"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if user, ok := auth.ParseJwtToken(authHeader); ok {
			c.Set("User", user)
			c.Set("IsAuthenticated", true)

		} else {
			c.Set("IsAuthenticated", false)

		}
		c.Next()
	}
}

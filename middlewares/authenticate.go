package middlewares

import (
	"github.com/akiratatsuhisa/first-golang-app/lib/auth"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if user, err := auth.ParseJwtToken(authHeader); err != nil {
			c.Set("IsAuthenticated", false)
		} else {
			c.Set("User", &user)
			c.Set("IsAuthenticated", true)
		}
		c.Next()
	}
}

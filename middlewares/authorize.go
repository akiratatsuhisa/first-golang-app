package middlewares

import (
	"net/http"

	"github.com/akiratatsuhisa/first-golang-app/auth"
	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, isAuthenticated := auth.GetUser(c); !isAuthenticated {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func AuthorizeRoles(roles []string) gin.HandlerFunc {
	hash := make(map[string]bool)

	for _, role := range roles {
		hash[role] = true
	}

	return func(c *gin.Context) {
		user, isAuthenticated := auth.GetUser(c)

		if !isAuthenticated {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		valid := false

		for _, role := range user.Roles {
			if hash[role] {
				valid = true
			}
		}

		if !valid {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}

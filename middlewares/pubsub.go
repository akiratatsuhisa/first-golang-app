package middlewares

import (
	"github.com/akiratatsuhisa/pubsub"
	"github.com/gin-gonic/gin"
)

func PubSub(ps *pubsub.MongoPubSub) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("PubSub", ps)

		c.Next()
	}
}

package middlewares

import (
	"github.com/akiratatsuhisa/pubsub"
	"github.com/gin-gonic/gin"
)

func PubSub(pubSub *pubsub.MongoPubSub) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("PubSub", pubSub)

		c.Next()
	}
}

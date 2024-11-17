package pubsub

import (
	"context"
	"os"

	"github.com/akiratatsuhisa/pubsub"
	"github.com/gin-gonic/gin"
)

func Initialize() *pubsub.MongoPubSub {
	ps := pubsub.NewMongoPubSub(&pubsub.MongoPubSubOpts{
		Ctx:        context.Background(),
		Uri:        os.Getenv("MONGODB_URI"),
		DbName:     "go_pubsub",
		CollName:   "pubsub",
		TtlSeconds: 900,
	})

	return ps
}

func GetPubSub(c *gin.Context) *pubsub.MongoPubSub {
	ps, _ := c.Get("PubSub")

	return ps.(*pubsub.MongoPubSub)
}

package controllers

import (
	"net/http"

	"github.com/akiratatsuhisa/first-golang-app/lib/pubsub"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type data struct {
	Message string `bson:"message" json:"message"`
}

func Send(c *gin.Context) {
	eventName := c.Param("eventName")

	ps := pubsub.GetPubSub(c)

	msg, _ := c.GetPostForm("message")

	ps.Publish(eventName, &data{
		Message: msg,
	})

	c.JSON(http.StatusOK, &gin.H{
		"message": "OK",
	})
}

func SSE(c *gin.Context) {
	eventName := c.Param("eventName")

	clientGone := c.Writer.CloseNotify()

	ps := pubsub.GetPubSub(c)

	listener := make(chan []byte)
	unsubscribe := ps.Subscribe(eventName, listener)
	defer unsubscribe()

	c.SSEvent("info", &gin.H{
		"message": "connected",
	})
	c.Writer.Flush()

Loop:
	for {
		select {
		case event := <-listener:
			var data data
			_ = bson.Unmarshal(event, &data)
			c.SSEvent("data", data)
			c.Writer.Flush()
		case <-clientGone:
			break Loop
		}
	}
}

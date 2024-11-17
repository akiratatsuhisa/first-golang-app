package main

import (
	"github.com/akiratatsuhisa/first-golang-app/db"
	"github.com/akiratatsuhisa/first-golang-app/lib/pubsub"
	"github.com/akiratatsuhisa/first-golang-app/middlewares"
	"github.com/akiratatsuhisa/first-golang-app/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	db.Connect()

	// Set up the router
	r := gin.Default()

	r.Use(cors.Default())

	r.Use(middlewares.Authenticate())

	r.Use(middlewares.PubSub(pubsub.Initialize()))

	router.Define(r)

	// Start the server
	r.Run("localhost:5000")
}

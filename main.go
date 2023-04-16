package main

import (
	"github.com/akiratatsuhisa/first-golang-app/db"
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

	router.Define(r)

	// Start the server
	r.Run(":5000")
}

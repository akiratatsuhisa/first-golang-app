package controllers

import (
	"net/http"

	"github.com/akiratatsuhisa/first-golang-app/db"
	"github.com/akiratatsuhisa/first-golang-app/models"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo

	if err := db.Context.Find(&todos).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &todos)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := db.Context.Create(&todo).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, &todo)
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo

	if err := db.Context.First(&todo, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, &todo)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo

	if err := db.Context.Where("id = ?", id).First(&todo).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Bind the request body to the todo model
	if err := c.BindJSON(&todo); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Update the todo in the database
	if err := db.Context.Save(&todo).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo

	if err := db.Context.Where("id = ?", id).First(&todo).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := db.Context.Delete(todo).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

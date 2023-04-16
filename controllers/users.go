package controllers

import (
	"net/http"

	"github.com/akiratatsuhisa/first-golang-app/auth"
	"github.com/akiratatsuhisa/first-golang-app/db"
	"github.com/akiratatsuhisa/first-golang-app/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	hash := auth.GenerateFromPassword(body.Password)

	user := models.User{
		Username: body.Username, Password: hash,
	}

	if err := db.Context.Create(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Username": user.Username,
	})
}

func Login(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var user models.User

	if err := db.Context.Where("username = ?", body.Username).Preload("UserRoles.Role").First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if !auth.CompareHashAndPassword(user.Password, body.Password) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateJwtToken(user)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"accessToken": token,
	})
}

func Profile(c *gin.Context) {
	user, _ := auth.GetUser(c)

	c.JSON(http.StatusOK, user)
}

package controllers

import (
	"database/sql"
	"net/http"

	"github.com/akiratatsuhisa/first-golang-app/db"
	"github.com/akiratatsuhisa/first-golang-app/lib/auth"
	"github.com/akiratatsuhisa/first-golang-app/lib/otp"
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

func GenerateTotpKey(c *gin.Context) {
	user, _ := auth.GetUser(c)

	uri, key := otp.GenerateUri(user.Username)

	c.JSON(http.StatusOK, gin.H{
		"Uri": uri,
		"Key": key,
	})
}

func SetTotpKey(c *gin.Context) {
	user, _ := auth.GetUser(c)

	var userInDB models.User

	if err := db.Context.Where("id = ?", user.ID).First(&userInDB).Error; err != nil || userInDB.TotpKey.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var body struct {
		Key string
	}

	if c.Bind(&body) != nil || body.Key == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Missing key",
		})
		return
	}

	userInDB.TotpKey = sql.NullString{String: body.Key, Valid: true}

	if err := db.Context.Save(&userInDB).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "OK",
	})
}

func ComapareTotp(c *gin.Context) {
	user, _ := auth.GetUser(c)

	var userInDB models.User

	if err := db.Context.Where("id = ?", user.ID).First(&userInDB).Error; err != nil || !userInDB.TotpKey.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var body struct {
		Otp string
	}

	if c.Bind(&body) != nil || body.Otp == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Missing otp",
		})
		return
	}

	message := "No"
	if otp.CompareOtp(userInDB.TotpKey.String, body.Otp) {
		message = "Yes"
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": message,
	})
}

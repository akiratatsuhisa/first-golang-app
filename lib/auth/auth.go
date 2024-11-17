package auth

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/akiratatsuhisa/first-golang-app/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		panic(err)
	}

	return string(hashed)
}

func CompareHashAndPassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))

	return err == nil
}

type User struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Roles    []string  `json:"roles"`
	Expires  time.Time `json:"expires"`
}

func GetUser(c *gin.Context) (*User, bool) {
	if isAuthenticated := c.GetBool("IsAuthenticated"); isAuthenticated {
		user, _ := c.Get("User")

		return user.(*User), isAuthenticated
	}

	return nil, false
}

var secret, _ = os.LookupEnv("SECRET")

func GenerateJwtToken(user *models.User) (string, error) {
	roles := []string{}
	for _, userRole := range user.UserRoles {
		roles = append(roles, userRole.Role.Name)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID,
		"username": user.Username,
		"roles":    roles,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	return claims.SignedString(([]byte(secret)))
}

func ParseJwtToken(authHeader string) (*User, error) {
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New("Not Found")
	}

	token, err := jwt.Parse(strings.TrimPrefix(authHeader, "Bearer "), func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("parse error")
	}

	if exp, err := claims.GetExpirationTime(); err != nil || exp.Before(time.Now()) {
		return nil, errors.New("jwt is expired")
	}

	roles := []string{}

	for _, value := range claims["roles"].([]interface{}) {
		if role, ok := value.(string); ok {
			roles = append(roles, role)
		}
	}

	var user = User{
		ID:       int(claims["sub"].(float64)),
		Username: claims["username"].(string),
		Roles:    roles,
		Expires:  time.Unix(int64(claims["exp"].(float64)), 0),
	}

	return &user, nil
}

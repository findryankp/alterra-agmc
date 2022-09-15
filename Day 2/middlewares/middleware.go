package middlewares

import (
	"fmt"
	"latihan/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	// "github.com/golang-jwt/jwt"
)

var jwtKey = []byte(os.Getenv("SECRETKEY"))

type JWTClaim struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	RoleID string `json:"role_id"`
	jwt.StandardClaims
}

func GenerateToken(user *models.User) (tokenString string, err error) {
	expirationTime := time.Now().Add(72 * time.Hour)
	fmt.Println(user.Email)
	claims := &JWTClaim{
		ID:     int(user.ID),
		Email:  user.Email,
		Name:   user.Name,
		RoleID: user.RoleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

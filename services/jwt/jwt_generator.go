// package jwt_services

// import (
// 	"os"

// 	"github.com/celpung/gocleanarch/entity"
// 	"github.com/golang-jwt/jwt"
// )

// type JwtService struct{}

// func NewJwtService() *JwtService {
// 	return &JwtService{}
// }

// func (js *JwtService) JWTGenerator(user entity.User) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"username": user.Username,
// 		"id":       user.ID,
// 		"role":     user.Role,
// 	})

// 	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN")))
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

package jwt_services

import (
	"os"
	"time"

	"github.com/celpung/gocleanarch/entity"
	"github.com/golang-jwt/jwt"
)

type JwtService struct{}

func NewJwtService() *JwtService {
	return &JwtService{}
}

func (js *JwtService) JWTGenerator(user entity.User) (string, error) {
	// Define the token expiration time
	expirationTime := time.Now().Add(24 * time.Hour).Unix() // Token valid for 24 hours

	// Create a new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"id":       user.ID,
		"role":     user.Role,
		"exp":      expirationTime,
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

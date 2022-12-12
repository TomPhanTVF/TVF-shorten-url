package security

import (
	model "user-service/internal/models"
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
)

const SecretKey = "fshjofjsdfo8oi3wyuf98wyu9876uhzxiou#@"

type JwtCustomClaims struct {
	Email 	string	`json:"email"`
	Role   	string	`json:"role"`
	jwt.StandardClaims
}

func Gentoken (user *model.User)(string, error){
	claims := &JwtCustomClaims{
		Email: user.GetEmail(),
		Role: user.GetRole(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour *24 ).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	resut, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err

	}
	return resut, nil
}

// Verify verifies the access token string and return a user claim if the token is valid
func Verify(accessToken string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(SecretKey), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
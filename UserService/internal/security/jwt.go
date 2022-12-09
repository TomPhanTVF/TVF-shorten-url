package security

import (
	model "user-service/internal/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SecretKey = "fshjofjsdfo8oi3wyuf98wyu9876uhzxiou#@"

type JwtCustomClaims struct {
	UserId 	string
	Email 	string
	Role   	string
	jwt.StandardClaims
}

func Gentoken (user model.User)(string, error){
	claims := &JwtCustomClaims{
		UserId: user.ID,
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


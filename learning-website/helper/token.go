package helper

import (
	"gin-socmed/model"
	"strconv"
	"time"

	"os"

	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte(os.Getenv("JWT_SECRET"))

type JWTClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(user *model.User) (string, error) {

	expiresIn, err := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN"))
	if err != nil {
		expiresIn = 3
	}

	claims := JWTClaims{
		ID: int(user.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiresIn) * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

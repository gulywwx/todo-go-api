package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gulywwx/todo-go-api/conf"
)

var jwtSecret = []byte(conf.JwtSetting.Secret)

type Claims struct {
	ID       string
	UserName string
	jwt.StandardClaims
}

/**
generate token
*/
func GeneratorToken(id string, username string) (string, error) {
	nowTime := time.Now()
	// expire time
	expireTime := nowTime.Add(conf.JwtSetting.ExpireTime * time.Hour)

	claims := Claims{
		ID:       id,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "todo",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

/**
parse token
*/
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

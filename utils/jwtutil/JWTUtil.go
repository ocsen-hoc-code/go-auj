package jwtutil

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ocsen-hoc-code/go-auj/models/user"
)

func CreateToken(user user.User, secretKey string, expiredTime int) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = user.ID
	atClaims["user_name"] = user.UserName
	atClaims["exp"] = time.Now().Add(time.Second * time.Duration(expiredTime)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidToken(token, secretKey string) (*user.User, bool) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	user := &user.User{}
	if err != nil {
		log.Println(err)
		return user, false
	}

	if !t.Valid {
		return user, false
	}

	var ok bool
	user.ID, ok = claims["user_id"].(string)
	if !ok {
		return user, false
	}
	user.UserName, ok = claims["user_name"].(string)
	if !ok {
		return user, false
	}

	return user, true
}

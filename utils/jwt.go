package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRET"))

func GenerateToken(email string) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email": email,
		"exp" : time.Now().Add(time.Hour*24).Unix(),
		"iat" : time.Now().Unix(),
	})

	tokenString,err := token.SignedString(secretKey)
	if err != nil {
		return "",err
	}

	return tokenString,nil
}

func VerifyToken(tokenString string) error{
	token,err := jwt.Parse(tokenString,func(token *jwt.Token) (interface{},error){
		return secretKey,nil
	})

	if err != nil{
		return err
	}

	if !token.Valid{
		return errors.New("invalid token")
	}

	return nil
}
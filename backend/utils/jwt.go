package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRET"))

func GenerateToken(email string,id int) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id" : id,
		"email": email,
		"exp" : time.Now().Add(time.Hour*1).Unix(),
		"iat" : time.Now().Unix(),
	})

	tokenString,err := token.SignedString(secretKey)
	if err != nil {
		return "",err
	}

	return tokenString,nil
}

func VerifyToken(tokenString string) (jwt.MapClaims,error){
	token,err := jwt.Parse(tokenString,func(token *jwt.Token) (interface{},error){
		return secretKey,nil
	})

	if err != nil{
		return nil,err
	}

	if !token.Valid{
		return nil,errors.New("invalid token")
	}

	claims,ok := token.Claims.(jwt.MapClaims)

	if !ok{
		return nil,errors.New("cannot parse claims")
	}	

	return claims,nil
}
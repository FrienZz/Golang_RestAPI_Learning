package middleware

import (
	"net/http"
	"strings"

	"github.com/FrienZz/Golang_RestAPI_Learning/utils"
	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context)  {
	s := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s,"Bearer")

	claims,err := utils.VerifyToken(token)

	if err != nil{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message" : "Unauthorized Access"})
		return
	}

	userId,ok := claims["user_id"].(float64)

	if !ok{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message" : "Unauthorized Access"})
		return
	}

	c.Set("userId",int(userId))
	c.Next()
}
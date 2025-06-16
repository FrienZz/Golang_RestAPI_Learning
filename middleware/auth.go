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

	err := utils.VerifyToken(token)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message" : "Unauthorized Access"})
		return
	}

	c.Next()
}
package handler

import (
	"net/http"

	"github.com/FrienZz/Golang_RestAPI_Learning/httphandle"
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	errRes,ok :=  err.(httphandle.ErrorResponse)
	if ok{
		c.JSON(errRes.Code, gin.H{"message" : errRes.Message})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"message" : "Unexpected error has occurred"})
}

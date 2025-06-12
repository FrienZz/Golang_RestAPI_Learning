package handler

import (
	"net/http"

	"github.com/FrienZz/Golang_RestAPI_Learning/models"
	"github.com/FrienZz/Golang_RestAPI_Learning/service"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *userHandler {
	return &userHandler{service: service}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var newUser models.User

	err := c.BindJSON(&newUser)
	
	if err != nil{
		HandleError(c,err)
		return
	}
	
	err = h.service.RegisterUser(newUser.Email,newUser.Password)

	if err != nil{
		HandleError(c,err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message" : "User has been successfully created!"})
}

func (h *userHandler) LoginUser(c *gin.Context) {

	var loginUser models.User
	
	err := c.BindJSON(&loginUser)
	
	if err != nil{
		HandleError(c,err)
		return
	}
	
	err = h.service.LoginUser(loginUser.Email,loginUser.Password)

	if err != nil{
		HandleError(c,err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Successfully login"})
}

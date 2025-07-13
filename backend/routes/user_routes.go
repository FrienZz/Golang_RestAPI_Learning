package routes

import (
	"database/sql"

	"github.com/FrienZz/Golang_RestAPI_Learning/external/adapter"
	"github.com/FrienZz/Golang_RestAPI_Learning/external/handler"
	"github.com/FrienZz/Golang_RestAPI_Learning/internal/service"
	"github.com/gin-gonic/gin"
)

func UserRoutes(db *sql.DB,server *gin.Engine){
	
	user_repo := adapter.NewUserRepositoryDB(db)
	user_service := service.NewUserService(user_repo)
	user_httpHandler := handler.NewUserHandler(user_service)
		
	server.POST("/register", user_httpHandler.RegisterUser)
	server.POST("/login", user_httpHandler.LoginUser)
}
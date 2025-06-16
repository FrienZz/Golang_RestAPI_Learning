package routes

import (
	"database/sql"

	"github.com/FrienZz/Golang_RestAPI_Learning/external/adapter"
	"github.com/FrienZz/Golang_RestAPI_Learning/external/handler"
	"github.com/FrienZz/Golang_RestAPI_Learning/internal/service"
	"github.com/FrienZz/Golang_RestAPI_Learning/middleware"
	"github.com/gin-gonic/gin"
)

func EventRoutes(db *sql.DB,server *gin.Engine){

	event_repo:= adapter.NewEventRepositoryDB(db)
	event_service := service.NewEventService(event_repo)
	event_httpHandler := handler.NewEventHandler(event_service)

	protected := server.Group("/",middleware.Authentication)
	
	protected.GET("/events",event_httpHandler.GetAllEvent)
	protected.GET("/events/:id", event_httpHandler.GetEvent)
	protected.POST("/events",event_httpHandler.CreateEvent)
	protected.PATCH("/events/:id",event_httpHandler.UpdateEvent)
	protected.DELETE("/events/:id",event_httpHandler.DeleteEvent)
	
}
package routes

import (
	"database/sql"

	"github.com/FrienZz/Golang_RestAPI_Learning/handler"
	"github.com/FrienZz/Golang_RestAPI_Learning/repository"
	"github.com/FrienZz/Golang_RestAPI_Learning/service"
	"github.com/gin-gonic/gin"
)

func EventRoutes(db *sql.DB,server *gin.Engine){

	event_repo:= repository.NewEventRepositoryDB(db)
	event_service := service.NewEventService(event_repo)
	event_httpHandler := handler.NewEventHandler(event_service)

	server.GET("/events", event_httpHandler.GetAllEvent)
	server.GET("/events/:id", event_httpHandler.GetEvent)
	server.POST("/events", event_httpHandler.CreateEvent)
	server.PATCH("/events/:id", event_httpHandler.UpdateEvent)
	server.DELETE("/events/:id", event_httpHandler.DeleteEvent)

}
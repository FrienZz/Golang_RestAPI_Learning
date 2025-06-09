package main

import (
	"log"

	"github.com/FrienZz/Golang_RestAPI_Learning/db"
	"github.com/FrienZz/Golang_RestAPI_Learning/handler"
	"github.com/FrienZz/Golang_RestAPI_Learning/repository"
	"github.com/FrienZz/Golang_RestAPI_Learning/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
 
  db := db.InitDB()
  defer db.Close()

  repo:= repository.NewEventRepositoryDB(db)
  service := service.NewEventService(repo)
  httpHandler := handler.NewEventHandler(service)
	
	router := gin.Default()

	router.GET("/events", httpHandler.GetAllEvent)
	router.GET("/events/:id", httpHandler.GetEvent)
	router.POST("/events", httpHandler.CreateEvent)
	router.PATCH("/events/:id", httpHandler.UpdateEvent)
	router.DELETE("/events/:id", httpHandler.DeleteEvent)
	router.Run(":8080")
}

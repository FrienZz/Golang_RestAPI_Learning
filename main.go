package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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

  psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",os.Getenv("PG_HOST"),  os.Getenv("PG_PORT"),  os.Getenv("PG_USER"),  os.Getenv("PG_PASSWORD"),  os.Getenv("PG_DB"))

	db,err := sql.Open("postgres",psqlInfo)

  if err != nil {
    log.Fatal(err)
  }

  defer db.Close()

  // Check the connection
  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Successfully connected!")
  repo:= repository.NewEventRepositoryDB(db)
  service := service.NewEventService(repo)
  httphandler := handler.NewEventHandler(service)

  router := gin.Default()

  router.GET("/events",httphandler.GetAllEvent)
  router.GET("/events/:id",httphandler.GetEvent)
  router.POST("/events",httphandler.CreateEvent)
  router.PATCH("/events/:id",httphandler.UpdateEvent)
  router.DELETE("/events/:id",httphandler.DeleteEvent)
  router.Run(":8080")
}

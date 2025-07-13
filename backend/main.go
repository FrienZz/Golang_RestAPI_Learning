package main

import (
	"log"
	"time"

	"github.com/FrienZz/Golang_RestAPI_Learning/db"
	"github.com/FrienZz/Golang_RestAPI_Learning/routes"
	"github.com/gin-contrib/cors"
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

  	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"}, 
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	
	routes.EventRoutes(db,server)
	routes.UserRoutes(db,server)
	
	server.Run(":8080")
	
}

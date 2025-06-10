package main

import (
	"log"

	"github.com/FrienZz/Golang_RestAPI_Learning/db"
	"github.com/FrienZz/Golang_RestAPI_Learning/routes"
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
	
	routes.EventRoutes(db,server)
	routes.UserRoutes(db,server)
	
	server.Run(":8080")
	
}

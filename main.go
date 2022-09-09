package main

import (
	"fmt"
	"log"

	"quiz/db"
	"quiz/middleware"
	"quiz/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	fmt.Println("Starting Go Backend Server")

	db.Init()

	router := createRouter()

	router.Run("quizbackend:3333")
}

func createRouter() *gin.Engine {
	router := gin.Default()

	public := router.Group("/")
	routes.PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private)

	return router
}

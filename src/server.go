package main

import (
	"go-gin-project/db"
	"go-gin-project/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func initServer() *gin.Engine {
	server := gin.Default()
	router.Setup(server)
	return server
}

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db.InitDB()
	defer db.CloseDB()

	server := initServer()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(server.Run(":" + port))
}

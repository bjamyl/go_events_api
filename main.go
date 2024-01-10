package main

import (
	"com.github/bjamyl/go-events-api/db"
	"com.github/bjamyl/go-events-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run(":8080")
}

// Getting all events

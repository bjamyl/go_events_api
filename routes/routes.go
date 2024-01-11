package routes

import (
	"com.github/bjamyl/go-events-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.POST("/events/:id", updateEvent)
	authenticated.POST("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}

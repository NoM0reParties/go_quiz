package routes

import (
	"github.com/gin-gonic/gin"

	"quiz/endpoint_handlers"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.POST("/user/create", eh.CreateUser)
	g.GET("/user/:id", eh.GetUser)
}
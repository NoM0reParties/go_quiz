package routes

import (
	"github.com/gin-gonic/gin"

	"quiz/endpoint_handlers"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.POST("/user/register", eh.CreateUser)
	g.POST("/user/login", eh.Login)
	g.GET("/user/:id", eh.GetUser)
}
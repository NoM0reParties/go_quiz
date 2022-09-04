package routes

import (
	"github.com/gin-gonic/gin"

	"quiz/endpoint_handlers"
)

func PrivateRoutes(g *gin.RouterGroup) {
	// Quiz Handlers
	g.GET("/quiz/list", eh.GetQuizList)
	g.GET("/quiz/:id", eh.GetQuiz)
	g.POST("/quiz/create", eh.CreateQuiz)
	g.PUT("/quiz/:id/update", eh.EditQuiz)
	g.DELETE("/quiz/:id/delete", eh.DeleteQuiz)
	// Theme Handlers
	g.GET("/theme/list", eh.GetThemeList)
	g.GET("/theme/:id", eh.GetTheme)
	g.POST("/theme/create", eh.CreateTheme)
	g.PUT("/theme/:id/update", eh.EditTheme)
	g.DELETE("/theme/:id/delete", eh.DeleteTheme)
}

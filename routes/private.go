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
	g.PATCH("/quiz/:id/update", eh.EditQuiz)
	g.DELETE("/quiz/:id/delete", eh.DeleteQuiz)
	g.GET("/quiz/:id/theme-info", eh.ThemeInfo)
	// Theme Handlers
	g.GET("/theme/list", eh.GetThemeList)
	g.GET("/theme/:id", eh.GetTheme)
	g.POST("/theme/create", eh.CreateTheme)
	g.PATCH("/theme/:id/update", eh.EditTheme)
	g.DELETE("/theme/:id/delete", eh.DeleteTheme)
	// Question Handlers
	g.GET("/question/list", eh.GetQuestionList)
	g.GET("/question/:id", eh.GetQuestion)
	g.POST("/question/create", eh.CreateQuestion)
	g.PUT("/question/:id/update", eh.EditQuestion)
	g.DELETE("/question/:id/delete", eh.DeleteQuestion)

	// MISC
	g.POST("/save-file", eh.SaveFileHandler)
}

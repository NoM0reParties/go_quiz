package eh

import (
	"net/http"
	"quiz/db"
	"quiz/dto"
	"quiz/helpers"

	"github.com/gin-gonic/gin"
)

func GetQuestionList(c *gin.Context) {
	var questions []db.Question
	var questionReposnse []dto.QuestionResponseDTO

	themeID := helpers.GetUint(c.Query("theme"))

	db.GetDB().Model(db.Question{}).Where("theme_id", themeID).Find(&questions)

	for _, question := range questions {
		questionReposnse = append(questionReposnse, question.GetDTO())
	}

	if questionReposnse == nil {
		c.IndentedJSON(http.StatusOK, []db.Question{})
	} else {
		c.IndentedJSON(http.StatusOK, questionReposnse)
	}

}

func CreateQuestion(c *gin.Context) {
	var question db.Question

	if err := c.BindJSON(&question); err != nil {
		c.AbortWithError(400, gin.Error{})
	}

	db.GetDB().Create(&question)

	c.IndentedJSON(http.StatusCreated, question.GetDTO())
}

func GetQuestion(c *gin.Context) {
	var question db.Question

	db.GetDB().Find(&question, c.Param("id"))

	if question.ID == 0 {
		c.IndentedJSON(400, dto.ErrorDetailDTO{Detail: "Instance not found"})
	} else {
		c.IndentedJSON(http.StatusOK, question.GetDTO())
	}
}

func EditQuestion(c *gin.Context) {
	var newQuestion db.Question
	var question db.Question

	if err := c.BindJSON(&newQuestion); err != nil {
		c.AbortWithError(400, gin.Error{})
	}

	db.GetDB().Find(&question, c.Param("id"))

	if question.ID == 0 {
		c.IndentedJSON(400, dto.ErrorDetailDTO{Detail: "Instance not found"})
	} else {
		question.TextContent = newQuestion.TextContent
		question.Value = newQuestion.Value
		question.MultimediaContent = newQuestion.MultimediaContent
		question.MultimediaType = newQuestion.MultimediaType

		db.GetDB().Model(&question).Where("id", c.Param("id")).Updates(question)

		c.IndentedJSON(http.StatusOK, question.GetDTO())
	}
}

func DeleteQuestion(c *gin.Context) {
	var Question db.Question

	db.GetDB().Delete(&Question, c.Param("id"))

	c.Writer.WriteHeader(204)
}
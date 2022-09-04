package eh

import (
	"net/http"
	"quiz/db"
	"quiz/dto"

	"github.com/gin-gonic/gin"
)

func GetQuizList(c *gin.Context) {
	var quizes []db.Quiz
	var quizReposnse []dto.QuizResponseDTO

	user, isUser := c.MustGet("user").(db.User)

	if !isUser {
		c.AbortWithError(400, gin.Error{})
	}

	db.GetDB().Model(db.Quiz{}).Where("user_id", user.ID).Find(&quizes)

	for _, quiz := range quizes {
		quizReposnse = append(quizReposnse, quiz.GetDTO())
	}

	c.IndentedJSON(http.StatusOK, quizReposnse)
}

func CreateQuiz(c *gin.Context) {
	var user db.User
	user, isUser := c.MustGet("user").(db.User)

	if !isUser {
		c.AbortWithError(400, gin.Error{})
	}

	var quiz db.Quiz

	if err := c.BindJSON(&quiz); err != nil {
		c.AbortWithError(400, gin.Error{})
	}

	quiz.UserID = user.ID
	quiz.User = user

	db.GetDB().Create(&quiz)

	c.IndentedJSON(http.StatusCreated, quiz.GetDTO())
}

func GetQuiz(c *gin.Context) {
	var quiz db.Quiz

	db.GetDB().Find(&quiz, c.Param("id"))

	if quiz.ID == 0 {
		c.IndentedJSON(400, dto.ErrorDetailDTO{Detail: "Instance not found"})
	} else {
		c.IndentedJSON(http.StatusOK, quiz.GetDTO())
	}
}

func EditQuiz(c *gin.Context) {
	var newQuiz db.Quiz
	var quiz db.Quiz

	if err := c.BindJSON(&newQuiz); err != nil {
		return
	}

	db.GetDB().Find(&quiz, c.Param("id"))

	if quiz.ID == 0 {
		c.IndentedJSON(400, dto.ErrorDetailDTO{Detail: "Instance not found"})
	} else {
		quiz.Completed = newQuiz.Completed
		quiz.SuperRound = newQuiz.SuperRound
		quiz.QuestionCount = newQuiz.QuestionCount
		quiz.ThemeCount = newQuiz.ThemeCount
		quiz.Name = newQuiz.Name

		db.GetDB().Model(&quiz).Where("id", c.Param("id")).Updates(quiz)

		c.IndentedJSON(http.StatusOK, quiz.GetDTO())
	}
}

func DeleteQuiz(c *gin.Context) {
	var quiz db.Quiz

	db.GetDB().Delete(&quiz, c.Param("id"))

	c.Writer.WriteHeader(204)
}

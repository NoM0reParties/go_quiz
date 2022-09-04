package eh

import (
	"net/http"
	"quiz/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUser(c *gin.Context) {
	var newUser db.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	newUser.Token = uuid.New()
	db.GetDB().Create(&newUser)

	c.IndentedJSON(http.StatusCreated, newUser)
}

func GetUser(c *gin.Context) {
	var newUser db.User

	db.GetDB().Find(&newUser, c.Param("id"))

	c.IndentedJSON(http.StatusCreated, newUser)
}


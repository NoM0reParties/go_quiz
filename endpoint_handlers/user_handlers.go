package eh

import (
	"net/http"
	"quiz/db"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var newUser db.User

	db.GetDB().Find(&newUser, c.Param("id"))

	c.IndentedJSON(http.StatusCreated, newUser)
}

package middleware

import (
	"quiz/db"
	"strings"

	"github.com/gin-gonic/gin"

	"log"
	"net/http"
)

func AuthRequired(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	var user db.User

	if strings.HasPrefix(token, "QAuth ") {
		token = strings.TrimPrefix(token, "QAuth ")
		db.GetDB().Where("token = ?", token).Find(&user)
		c.Set("user", user)
	}

	if user.ID == 0 {
		log.Println("User not logged in")
		c.AbortWithStatus(http.StatusForbidden)
	}
	c.Next()
}
package eh

import (
	"fmt"
	"quiz/db"
	"quiz/dto"
	"quiz/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var newUser db.User

	if err := c.BindJSON(&newUser); err != nil {
		c.AbortWithError(400, gin.Error{})
	}

	newUser.Password = helpers.HashAndSalt([]byte(newUser.Password))

	newUser.Token = uuid.New()
	db.GetDB().Create(&newUser)

	c.IndentedJSON(201, newUser)
}

func Login(c *gin.Context) {
	var userData dto.LoginDTO
	var user db.User

	if err := c.BindJSON(&userData); err != nil {
		c.AbortWithError(400, gin.Error{})
	}

	db.GetDB().Where(&db.User{Name: userData.Name}).First(&user)
	fmt.Println(userData.Name)

	encryptionErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))

	if encryptionErr != nil {
		c.IndentedJSON(403, dto.ErrorDetailDTO{Detail: "Login or Password are incorrect"})
	} else {
		c.IndentedJSON(201, &user)
	}
}

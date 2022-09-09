package eh

import (
	"fmt"
	"path/filepath"
	"quiz/db"
	"quiz/dto"
	"quiz/helpers"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
    _, b, _, _ = runtime.Caller(0)
    basepath   = filepath.Dir(b)
)

func CreateUser(c *gin.Context) {
	var regDTO dto.RegisterDTO

	if err := c.ShouldBind(&regDTO); err != nil {
		c.AbortWithError(400, err)
	}

	fmt.Println(basepath)
	newUser := db.User{
		Name:  regDTO.Name,
	}

	if regDTO.File != nil {
		extension := filepath.Ext(regDTO.File.Filename)
		fmt.Println(extension)
		file := regDTO.File
		dst := "/media/" + uuid.New().String() + extension

		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.AbortWithError(400, err)
		}

		fmt.Println(extension)
		newUser.Photo = dst
	}

	newUser.Password = helpers.HashAndSalt([]byte(regDTO.Password))

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

	encryptionErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))

	if encryptionErr != nil {
		c.IndentedJSON(403, dto.ErrorDetailDTO{Detail: "Login or Password are incorrect"})
	} else {
		c.IndentedJSON(201, &user)
	}
}

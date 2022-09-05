package eh

import (
	"fmt"
	"path/filepath"
	"quiz/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveFileHandler(c *gin.Context) {
	var bindFile dto.FileDTO

	// Bind file
	if err := c.ShouldBind(&bindFile); err != nil {
		c.String(400, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	// Save uploaded file
	extension := filepath.Ext(bindFile.File.Filename)
	file := bindFile.File
	dst := "/home/fedor/Documents/pet_projects/go_quiz/media/" + uuid.New().String() + extension
	
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.String(400, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
}
package eh

import (
	"net/http"
	"quiz/db"
	"quiz/dto"
	"quiz/helpers"

	"github.com/gin-gonic/gin"
)

func GetThemeList(c *gin.Context) {
	var themes []db.Theme
	var themeReposnse []dto.ThemeResponseDTO

	quizID := helpers.GetUint(c.Query("quiz"))
	
	db.GetDB().Model(db.Theme{}).Where("quiz_id", quizID).Find(&themes)

	for _, theme := range themes {
		themeReposnse = append(themeReposnse, theme.GetDTO())
	}

	if themeReposnse == nil {
		c.IndentedJSON(http.StatusOK, []db.Theme{})
	} else {
		c.IndentedJSON(http.StatusOK, themeReposnse)
	}
	
}

func CreateTheme(c *gin.Context) {
	var theme db.Theme
	
	if err := c.BindJSON(&theme); err != nil {
		c.AbortWithError(400, gin.Error{})
	}
	
	db.GetDB().Create(&theme)

	c.IndentedJSON(http.StatusCreated, theme.GetDTO())
}

func GetTheme(c *gin.Context) {
	var theme db.Theme

	db.GetDB().Find(&theme, c.Param("id"))

	if theme.ID == 0 {
		c.IndentedJSON(400, dto.ErrorDetailDTO{Detail: "Instance not found"})
	} else {
		c.IndentedJSON(http.StatusOK, theme.GetDTO())
	}
}

func EditTheme(c *gin.Context) {
	var newTheme db.Theme
	var theme db.Theme

	if err := c.BindJSON(&newTheme); err != nil {
		c.AbortWithError(400, gin.Error{})
	}

	db.GetDB().Find(&theme, c.Param("id"))

	if theme.ID == 0 {
		c.IndentedJSON(400, dto.ErrorDetailDTO{Detail: "Instance not found"})
	} else {
		theme.Name = newTheme.Name

		db.GetDB().Model(&theme).Where("id", c.Param("id")).Updates(theme)

		c.IndentedJSON(http.StatusOK, theme.GetDTO())
	}
}

func DeleteTheme(c *gin.Context) {
	var theme db.Theme

	db.GetDB().Delete(&theme, c.Param("id"))

	c.Writer.WriteHeader(204)
}

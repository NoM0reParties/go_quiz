package dto

import "mime/multipart"

type FileDTO struct {
	File *multipart.FileHeader `form:"file"`
}

type UserDTO struct {
	Name string
}

type QuizResponseDTO struct {
	ID            uint
	Name          string
	Completed     bool
	QuestionCount int
	ThemeCount    int
	SuperRound    bool
}

type ThemeResponseDTO struct {
	ID    uint
	Name  string
	Round int
}

type QuestionResponseDTO struct {
	ID                uint
	TextContent       string
	Value             int32
	MultimediaContent string
	MultimediaType    string
}

type LoginDTO struct {
	Name     string
	Password string
}

type RegisterDTO struct {
	Name     string                `form:"name" binding:"required"`
	Password string                `form:"password"`
	File     *multipart.FileHeader `form:"file"`
}

type ThemeInfoDTO struct {
	ThemesAvailable []int32
	ThemeCount      int32
}

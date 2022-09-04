package dto

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
	ID            uint
	Name          string
}
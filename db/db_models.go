package db

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Photo string
}

type Achivement struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	Level  string
	UserID uint
	User   User `gorm:"foreignKey:UserID"`
}

type Quiz struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Completed bool
}

type Theme struct {
	ID   uint
	Name string
}

type Question struct {
	ID                uint `gorm:"primaryKey"`
	TextContent       string
	Value             int32
	MultimediaContent string
}

type InGameQuestion struct {
	ID         uint `gorm:"primaryKey"`
	QuestionID uint
	Question   Question `gorm:"foreignKey:QuestionID"`
	IsAnswered bool
}

type Participant struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	User   User `gorm:"foreignKey:UserID"`
	GameID uint
	Game   Game `gorm:"foreignKey:GameID"`
	Score  int32
}

type Game struct {
	ID           uint `gorm:"primaryKey"`
	QuizID       uint
	Quiz         Quiz `gorm:"foreignKey:QuizID"`
	WinnerID     uint
	Winner       User `gorm:"foreignKey:WinnerID"`
	Participants []Participant
}


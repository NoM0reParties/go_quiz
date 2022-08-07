package db

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
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
	Question   Question
	IsAnswered bool
}

type Game struct {
	ID           uint `gorm:"primaryKey"`
	Quiz         Quiz
	Winner       User
	Participants []User
}

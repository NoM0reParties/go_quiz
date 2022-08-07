package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	pgUser, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		log.Fatal("No enviroment wasn't set")
	}
	pgPass, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {
		log.Fatal("No enviroment wasn't set")
	}
	pgDB, exists := os.LookupEnv("POSTGRES_DB")
	if !exists {
		log.Fatal("No enviroment wasn't set")
	}
	pgPort, exists := os.LookupEnv("POSTGRES_PORT")
	if !exists {
		log.Fatal("No enviroment wasn't set")
	}
	pgHost, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		log.Fatal("No enviroment wasn't set")
	}

	dbURL := "postgres://" + pgUser + ":" + pgPass + "@" + pgHost + ":" + pgPort + "/" + pgDB

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(
		&User{},
		&Quiz{},
		&Theme{},
		&Question{},
		&InGameQuestion{},
		&Game{},
	)

	return db
}

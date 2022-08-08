package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBCLIENT *gorm.DB
var err error

func Init() {

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

	dbURL := " host=" + pgHost + " user=" + pgUser + " password=" + pgPass + " dbname=" + pgDB + " port=" + pgPort + " sslmode=disable TimeZone=UTC"

	// dbURL := "postgres://" + pgUser + ":" + pgPass + "@" + pgHost + ":" + pgPort + "/" + pgDB

	DBCLIENT, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully connected to DB")

	DBCLIENT.AutoMigrate(
		&User{},
		&Quiz{},
		&Theme{},
		&Question{},
		&InGameQuestion{},
		&Game{},
		&Participant{},
	)
}

func GetDB() *gorm.DB {
	return DBCLIENT
}

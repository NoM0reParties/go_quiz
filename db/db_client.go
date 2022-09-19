package db

import (
	"fmt"
	"log"
	"os"
	"time"

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

	for i := 1; i <= 10; i++ {
		DBCLIENT, err = gorm.Open(postgres.New(postgres.Config{
			DSN: dbURL,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
		if err == nil {
			i = 10
		} else {
			fmt.Println(dbURL)
			time.Sleep(3 * time.Second)
			fmt.Println(err)
		}
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
		&Achivement{},
	)
}

func GetDB() *gorm.DB {
	return DBCLIENT
}

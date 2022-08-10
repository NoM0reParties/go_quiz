package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"quiz/db"
	"quiz/endpoint_handlers"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	db.Init()

	router := createRouter()
	http.Handle("/", router)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func createRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// User endpoints
	router.HandleFunc("/user/create", eh.CreateUser).Methods("POST")

	return router
}

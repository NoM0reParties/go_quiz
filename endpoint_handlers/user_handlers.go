package eh

import (
	"encoding/json"
	"net/http"
	"quiz/db"
)


func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	
	var newUser db.User

	json.NewDecoder(r.Body).Decode(&newUser)

	db.GetDB().Create(&newUser)

	json.NewEncoder(w).Encode(newUser)
}


package eh

import (
	"encoding/json"
	"net/http"
	"quiz/db"
)

func CreateQuiz(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	
	var newQuiz db.Quiz

	json.NewDecoder(r.Body).Decode(&newQuiz)

	db.GetDB().Create(&newQuiz)

	json.NewEncoder(w).Encode(newQuiz)
}

func EditQuiz(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	
	var quiz db.Quiz

	var newQuiz db.Quiz

	dbClient := db.GetDB()

	json.NewDecoder(r.Body).Decode(&newQuiz)

	dbClient.Find(&quiz, newQuiz.ID)

	quiz.Name = newQuiz.Name
	quiz.Completed = newQuiz.Completed

	dbClient.Save(&quiz)

	json.NewEncoder(w).Encode(newQuiz)
}

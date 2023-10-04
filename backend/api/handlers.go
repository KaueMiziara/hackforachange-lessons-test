package api

import (
	"encoding/json"
	"net/http"
)

func GetGrades(w http.ResponseWriter, r *http.Request) {
	grades, err := LoadGrades()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(grades)
}

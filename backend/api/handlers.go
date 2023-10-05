package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func GetSubjectByGrade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gradeID, err := strconv.Atoi(vars["gradeID"])
	if err != nil {
		http.Error(w, "Invalid grade ID", http.StatusBadRequest)
		return
	}

	subjects, err := LoadSubjectByGradeId(gradeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subjects)
}

func GetLessonsByGradeAndSubject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gradeID, err := strconv.Atoi(vars["gradeID"])
	if err != nil {
		http.Error(w, "Invalid grade ID", http.StatusBadRequest)
		return
	}

	subjectID, err := strconv.Atoi(vars["subjectID"])
	if err != nil {
		http.Error(w, "Invalud subject ID", http.StatusBadRequest)
		return
	}

	lessons, err := LoadLessonsByGradeAndSubject(gradeID, subjectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lessons)
}

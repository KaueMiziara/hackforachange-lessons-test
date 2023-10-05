package api

import "github.com/gorilla/mux"

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/api/grades", GetGrades).Methods("GET")
	router.HandleFunc("/api/subjects/{gradeID}", GetSubjectByGrade).Methods("GET")
	router.HandleFunc("/api/lessons/{gradeID}/{subjectID}", GetLessonsByGradeAndSubject).Methods("GET")
}

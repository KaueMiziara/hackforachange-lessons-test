package api

import "github.com/gorilla/mux"

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/api/grades", GetGrades).Methods("GET")
}

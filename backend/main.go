package main

import (
	"hackforachange-demo-backend/api"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	api.SetupRoutes(router)

	http.ListenAndServe(":8080", router)
}

package main

import (
	"hackforachange-swift-backend/api"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	api.SetupRoutes(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
}

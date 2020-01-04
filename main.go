package main

import (
	"fuzzy-umbrella/app"
	"fuzzy-umbrella/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Info).Methods("GET")

	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Authenticate).Methods("POST")
	router.HandleFunc("/account", handlers.GetUser).Methods("GET")

	router.HandleFunc("/graphql", handlers.GraphQL).Methods("POST")

	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("Listning at http://localhost:" + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}

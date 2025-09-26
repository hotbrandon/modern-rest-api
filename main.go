package main

import (
	"fmt"
	"hotbrandon/modern-rest-api/internal/handler"
	"net/http"
)

func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
		next(w, r)
	}
}

func setupRoutes() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "alive!")
	})

	http.HandleFunc("POST /api/v1/auth/login", logRequest(handler.HandleLogin))
	// create a new user
	http.HandleFunc("POST /api/v1/users", logRequest(handler.HandleCreateUser))
	// get all users
	http.HandleFunc("GET /api/v1/users", logRequest(handler.HandleGetUsers))

}

func main() {
	setupRoutes()

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

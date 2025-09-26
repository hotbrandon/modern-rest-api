package main

import (
	"fmt"
	"hotbrandon/modern-rest-api/internal/handler"
	"hotbrandon/modern-rest-api/internal/middleware"
	"net/http"
)

func setupRoutes() {
	http.HandleFunc("/health", middleware.RequireAuth(checkHealth))

	http.HandleFunc("POST /api/v1/auth/login", middleware.LogRequest(handler.HandleLogin))
	// create a new user
	http.HandleFunc("POST /api/v1/users", middleware.LogRequest(handler.HandleCreateUser))
	// get all users
	http.HandleFunc("GET /api/v1/users", middleware.LogRequest(handler.HandleGetUsers))

}

func checkHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "alive!")
}

func main() {
	setupRoutes()

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

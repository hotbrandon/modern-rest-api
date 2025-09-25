package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type Payload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username string `json:"username"`
}

func randNum(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(50)
	fmt.Fprint(w, num)

}

func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
		next(w, r)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	payload := Payload{}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := User{Username: payload.Username}

	// Encode to buffer first to catch encoding errors before writing headers
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func getUsers(w http.ResponseWriter, _ *http.Request) {
	users := []User{
		{Username: "user1"},
		{Username: "user2"},
		{Username: "user3"},
	}

	// Encode to buffer first to catch encoding errors before writing headers
	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		fmt.Println("encode error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createUser(w, r)
	case http.MethodGet:
		getUsers(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

	}

}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	http.HandleFunc("/random", randNum)
	http.HandleFunc("/users", logRequest(handleUsers)) // create a new user
	http.HandleFunc("/lists", logRequest(handleList))  // create a new user

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

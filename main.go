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

func randNum(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(50)
	fmt.Fprint(w, num)

}

func handleUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintln(w, "Create a new user")
		payload := Payload{}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Example: respond with the parsed values
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"user":     payload.Username,
			"password": payload.Password,
		})
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	http.HandleFunc("/random", randNum)

	http.HandleFunc("/users", handleUser)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

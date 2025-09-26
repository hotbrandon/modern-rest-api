package handler

import (
	"encoding/json"
	"hotbrandon/modern-rest-api/internal/util"
	"log"
	"net/http"
	"time"
)

var users []User = []User{}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type Session struct {
	Username string
	Expire   time.Time
}

var sessions map[string]Session = make(map[string]Session)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	payload := LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.Username == payload.Username && user.Password == payload.Password {
			token, err := util.GeterateJwtToken(user.Username)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			response := LoginResponse{
				Token: token,
			}
			jsonData, err := json.Marshal(response)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			session := Session{
				Username: user.Username,
				Expire:   time.Now().Add(time.Hour * 24),
			}
			sessions[token] = session
			log.Printf("Sessions: %v\n", sessions)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
			return
		}
	}
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	payload := LoginRequest{}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := User{
		ID:       len(users) + 1,
		Username: payload.Username,
		Password: payload.Password,
	}

	// Encode to buffer first to catch encoding errors before writing headers
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// TODO persist the new user to the database
	users = append(users, user)
	log.Printf("New user created: %s\n", user.Username)
	log.Printf("HandleCreateUser: %v\n", users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func HandleGetUsers(w http.ResponseWriter, _ *http.Request) {
	// Encode to buffer first to catch encoding errors before writing headers
	log.Printf("HandleGetUsers: %v\n", users)
	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func GetSession(token string) (Session, bool) {
	session, ok := sessions[token]
	return session, ok
}

func GetUser(userName string) *User {
	for _, user := range users {
		if user.Username == userName {
			return &user
		}
	}
	return nil
}

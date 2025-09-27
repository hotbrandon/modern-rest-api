package handler

import (
	"encoding/json"
	"hotbrandon/modern-rest-api/internal/service"
	"log"
	"net/http"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: svc}

}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	payload := LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(payload.Username, payload.Password)
	if err != nil {
		http.Error(w, "login failed", http.StatusUnauthorized)
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
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (h *AuthHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	payload := LoginRequest{}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.authService.CreateUser(payload.Role, payload.Username, payload.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *AuthHandler) HandleGetUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := h.authService.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

// func (h *AuthHandler) GetSession(token string) (Session, bool) {
// 	session, ok := sessions[token]
// 	return session, ok
// }

// func (h *AuthHandler) GetUser(userName string) *User {
// 	for _, user := range users {
// 		if user.Username == userName {
// 			return &user
// 		}
// 	}
// 	return nil
// }

package main

import (
	"database/sql"
	"fmt"
	"hotbrandon/modern-rest-api/internal/handler"
	"hotbrandon/modern-rest-api/internal/middleware"
	"hotbrandon/modern-rest-api/internal/repository"
	"hotbrandon/modern-rest-api/internal/service"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func setupRoutes(authHandler *handler.AuthHandler) {
	http.HandleFunc("/health", checkHealth)

	http.HandleFunc("POST /api/v1/auth/login", middleware.LogRequest(authHandler.HandleLogin))
	// create a new user
	http.HandleFunc("POST /api/v1/users", middleware.RequireAdmin(middleware.LogRequest(authHandler.HandleCreateUser)))
	// get all users
	http.HandleFunc("GET /api/v1/users", middleware.RequireAdmin(middleware.LogRequest(authHandler.HandleGetUsers)))

}

func checkHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "alive!")
}

func main() {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Configure pool settings
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(1 * time.Hour)

	if err := db.Ping(); err != nil {
		log.Println(err)
		return
	}

	repo := repository.NewRepository(db)
	if err := repo.Init(); err != nil {
		log.Println(err)
		return
	}

	authService := service.NewAuthService(repo)
	authHandler := handler.NewAuthHandler(authService)

	setupRoutes(authHandler)

	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

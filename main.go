package main

import (
	"database/sql"
	"fmt"
	"hotbrandon/modern-rest-api/internal/handler"
	"hotbrandon/modern-rest-api/internal/middleware"
	"hotbrandon/modern-rest-api/internal/repository"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func setupRoutes(_ *handler.Handler) {
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
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Configure pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(60 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	if err := repo.Init(); err != nil {
		log.Fatal(err)
	}

	handler := handler.NewHandler(repo)

	setupRoutes(handler)

	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

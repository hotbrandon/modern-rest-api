package repository

import (
	"database/sql"
	"hotbrandon/modern-rest-api/internal/model"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Init() error {
	sql := `CREATE TABLE IF NOT EXISTS users
			(
				role VARCHAR,
				username VARCHAR PRIMARY KEY, 
				password VARCHAR
			)`
	if _, err := r.db.Exec(sql); err != nil {
		return err
	}

	sql = `CREATE TABLE IF NOT EXISTS sessions 
			(
				token VARCHAR PRIMARY KEY, 
				expires TIMESTAMP, 
				username VARCHAR
			)`
	if _, err := r.db.Exec(sql); err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateUser(role, username, password string) error {
	count, err := r.GetUserByName(username)

	if count > 0 {
		return err
	}

	sql := `INSERT INTO users (role, username, password) VALUES (?, ?, ?)`
	_, err = r.db.Exec(sql, role, username, password)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUsers() ([]model.User, error) {
	sql := `SELECT role, username FROM users`
	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Role, &user.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *Repository) GetUserByName(username string) (int, error) {
	sql := `SELECT count(*) FROM users WHERE username = ?`
	count := 0
	err := r.db.QueryRow(sql, username).Scan(&count)

	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repository) ValidateUser(username, password string) (bool, error) {
	sql := `SELECT username FROM users WHERE username = ? AND password = ?`
	row := r.db.QueryRow(sql, username, password)
	user := model.User{}
	err := row.Scan(&user.Username)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *Repository) CreateSession(token string, expires time.Time, username string) error {
	sql := `INSERT INTO sessions (token, expires, username) VALUES (?, ?, ?)`
	_, err := r.db.Exec(sql, token, expires, username)
	return err
}

func (r *Repository) GetSession(token string) (*model.Session, error) {
	sql := `SELECT token, expires, username FROM sessions WHERE token = ?`
	row := r.db.QueryRow(sql, token)
	session := model.Session{}
	err := row.Scan(&session.Token, &session.Expire, &session.Username)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

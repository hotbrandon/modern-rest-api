package repository

import (
	"database/sql"

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

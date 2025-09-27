package model

import (
	"time"
)

type User struct {
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type Session struct {
	Token    string
	Username string
	Expire   time.Time
}

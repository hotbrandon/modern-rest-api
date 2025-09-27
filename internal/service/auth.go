package service

import (
	models "hotbrandon/modern-rest-api/internal/model"
	"hotbrandon/modern-rest-api/internal/repository"
	"hotbrandon/modern-rest-api/internal/util"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(role, username, password string) error {
	err := s.repo.CreateUser(role, username, password)
	return err
}

func (s *AuthService) GetUsers() ([]models.User, error) {
	return s.repo.GetUsers()
}

func (s *AuthService) Login(username, password string) (string, error) {
	_, err := s.repo.ValidateUser(username, password)
	if err != nil {
		return "", err
	}

	token, err := util.GeterateJwtToken(username)
	if err != nil {
		return "", err
	}

	err = s.repo.CreateSession(token, time.Now().Add(time.Hour*24), username)
	if err != nil {
		return "", err
	}

	return token, nil
}

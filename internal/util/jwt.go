package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GeterateJwtToken(userId string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := []byte("hmacSampleSecret")
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func VerifyToken(token string) (*jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("hmacSampleSecret"), nil
	})
	if err != nil {
		return nil, err
	}
	return &claims, nil
}

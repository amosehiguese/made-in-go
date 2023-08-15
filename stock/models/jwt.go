package models

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type Token struct {
	Access  string
	Refresh string
}

func GenerateNewTokens(id string, permissions []string) (*Token, error) {
	accessToken, err := newAccessToken(id, permissions)
	if err != nil {
		log.Println("Error creating access token ->", err)
		return nil, err
	}

	refreshToken, err := newRefreshToken(id, permissions)
	if err != nil {
		log.Println("Error creating refresh token ->", err)
		return nil, err
	}

	return &Token{
		Access: accessToken,
		Refresh: refreshToken,
	}, nil
}

func newAccessToken(id string, perms []string) (string, error) {
	AccessTokenSecret := os.Getenv("JWT_ACCESS_SECRET_KEY")
	expires, err := strconv.Atoi(os.Getenv("JWT_TOKEN_EXPIRY_TIME"))
	if err != nil {
		log.Println("Got an error while trying to convert JWT_TOKEN_EXPIRY_TIME from a string to an int ->", err)
	}

	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["expires"] = time.Now().Add(time.Duration(expires) * time.Minute).Unix()

	for _, perm := range perms {
		claims[perm] = true
	}
}

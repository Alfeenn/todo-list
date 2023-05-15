package model

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var Key = os.Getenv("JWT_SECRET_KEY")

type Token struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Session struct {
	Id       string    `json:"id"`
	Username string    `json:"username"`
	Token    string    `json:"token"`
	Expiry   time.Time `json:"exp"`
}

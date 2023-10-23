package model

import (
	"gorm.io/gorm"
	"os"
)

type GithubUser struct {
	gorm.DB
	GithubUserID int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Role         string `json:"role"`
}

var JwtKey = []byte(os.Getenv("JWT_KEY"))

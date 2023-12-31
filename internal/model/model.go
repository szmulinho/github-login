package model

import (
	"database/sql/driver"
	"encoding/json"
	"os"
)

type GhUser struct {
	Login       string `json:"login" gorm:"index"`
	AvatarUrl   string `json:"avatar_url" gorm:"index"`
	Role        string `json:"role" gorm:"index"`
}

type PublicRepo struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
}

func (u *GhUser) Value() (driver.Value, error) {
	return json.Marshal(u)
}

func (u *GhUser) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), u)
}

type LoginResponse struct {
	User  GhUser   `json:"user"`
	Token string `json:"token"`
}

type LoginResponse2 struct {
	Doctor  GhUser   `json:"doctor"`
	Token string `json:"token"`
}

var JwtKey = []byte(os.Getenv("JWT_KEY"))



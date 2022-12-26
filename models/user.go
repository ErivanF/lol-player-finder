package models

import (
	"time"
)

type UserCreateRequest struct {
	Name       string `json:"name"`
	InGameName string `json:"inGameName"`
	Password   string `json:"password"`
	Email      string `json:"email"`
}
type UserDB struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	InGameName string    `json:"inGameName"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"createdcreatedAt"`
}

package types

import (
	"time"
)

type UserCreate struct {
	Name     string `json:"name" sql:"name"`
	Email    string `json:"email" sql:"email" `
	Password string `json:"password" sql:"password" `
	GameName string `json:"gameName" sql:"gameName" `
}
type User struct {
	Id        int       `json:"id" sql:"id"`
	Name      string    `json:"name" sql:"name"`
	Email     string    `json:"email" sql:"email" `
	GameName  string    `json:"gameName" sql:"gameName" `
	CreatedAt time.Time `json:"createdAt" sql:"created_at" `
	UpdatedAt time.Time `json:"updatedAt" sql:"updated_at" `
}
type UserChanges struct {
	Name     string `json:"name" sql:"name"`
	Password string `json:"password" sql:"password" `
	GameName string `json:"gameName" sql:"gam_name" `
}

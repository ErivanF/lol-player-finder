package user_controller

import (
	"fmt"
	"main/db"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Create(c *gin.Context) {
	var new_user models.UserCreateRequest
	err := c.BindJSON(&new_user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	if new_user.Name == "" {
		panic("Name required")
	}
	hashed_password, _ := bcrypt.GenerateFromPassword([]byte(new_user.Password), 8)
	new_user.Password = string(hashed_password)
	stmt, err := db.Conn.Prepare(`
		INSERT INTO users
		(name, in_game_name, password, email)
		VALUES
		($1, $2, $3, $4)
		RETURNING *
		`)
	result := stmt.QueryRow(
		new_user.Name,
		new_user.InGameName,
		new_user.Password,
		new_user.Email)

	var created_user models.UserDB

	if err = result.Scan(
		&created_user.ID,
		&created_user.Name,
		&created_user.InGameName,
		&created_user.Password,
		&created_user.Email,
		&created_user.CreatedAt); err != nil {
		fmt.Println("Erro: ", err)
	}

	created_user.Password = ""
	c.IndentedJSON(http.StatusOK, created_user)
	c.IndentedJSON(http.StatusOK, new_user)
}

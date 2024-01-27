package userServices

import (
	"fmt"
	"lol-player-finder/database"
	"lol-player-finder/types"
)

func CreateUserService(user types.UserCreate) types.User {
	conn := database.GedDb()
	query := `	INSERT INTO users 
				(name, email, password, gameName)
				VALUES ($1, $2, $3, $4)
				RETURNING id, name, email, gameName, created_at, updated_at`
	var newUser types.User
	err := conn.QueryRow(query,
		user.Name,
		user.Email,
		user.Password,
		user.GameName).Scan(
		&newUser.Id,
		&newUser.Name,
		&newUser.Email,
		&newUser.GameName,
		&newUser.CreatedAt,
		&newUser.UpdatedAt)

	if err != nil {
		fmt.Println(err.Error())
	}

	return newUser

}

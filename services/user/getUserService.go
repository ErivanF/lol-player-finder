package userServices

import (
	"fmt"
	"lol-player-finder/database"
	"lol-player-finder/types"
)

func GetUserService(id int64) types.User {
	query := `
				SELECT 
				id, name, email, gameName, created_at, updated_at
				FROM users
				where id = $1
	`
	conn := database.GetDb()
	var user types.User
	err := conn.QueryRow(query, id).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.GameName,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		fmt.Println(err.Error())
	}
	return user
}

package userServices

import (
	"fmt"
	"lol-player-finder/database"
	"lol-player-finder/types"
)

func GetUsersService() []types.User {
	var users []types.User
	conn := database.GetDb()
	query := `SELECT id, name, email, gameName, created_at, updated_at FROM users`
	rows, err := conn.Query(query)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		var user types.User
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.GameName,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			fmt.Println(err.Error())
		}
		users = append(users, user)
	}
	return users
}

package userServices

import (
	"fmt"
	"lol-player-finder/database"
	appError "lol-player-finder/error"
	"lol-player-finder/types"
)

func CreateUserService(user types.UserCreate) (types.User, error) {
	conn := database.GetDb()
	query := `	INSERT INTO users 
				(name, email, password, game_name)
				VALUES ($1, $2, $3, $4)
				RETURNING id, name, email, game_name, created_at, updated_at`
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
		return types.User{}, appError.InternalServerError("database connection error")
	}

	return newUser, nil

}

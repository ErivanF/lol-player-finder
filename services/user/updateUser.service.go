package userServices

import (
	"lol-player-finder/database"
	appError "lol-player-finder/error"
	"lol-player-finder/types"
)

func UpdateUserService(id string, changes types.UserChanges) (types.User, error) {
	query := `
	    SELECT name, password, game_name
		FROM users WHERE id =$1`

	var oldUser types.UserChanges
	conn := database.GetDb()
	err := conn.QueryRow(query, id).Scan(&oldUser.Name, &oldUser.Password, &oldUser.GameName)
	if err != nil {
		return types.User{}, appError.BadRequestError("Invalid ID or database connection")
	}
	if changes.Name != "" {
		oldUser.Name = changes.Name
	}
	if changes.Password != "" {
		oldUser.Password = changes.Password
	}
	if changes.GameName != "" {
		oldUser.GameName = changes.GameName
	}

	query = `
			UPDATE users SET
			name = $1, 
			password = $2, 
			game_name = $3,
			updated_at = CURRENT_TIMESTAMP
			WHERE id = $4
			RETURNING
			id,
			email,
			created_at, 
			updated_at
	`
	var user types.User
	err = conn.QueryRow(
		query,
		oldUser.Name,
		oldUser.Password,
		oldUser.GameName,
		id).Scan(
		&user.Id,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	user.Name = oldUser.Name
	user.GameName = oldUser.GameName
	return user, nil
}

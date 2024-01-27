package userServices

import "lol-player-finder/database"

func IsEmailAvailableService(email string) bool {
	conn := database.GedDb()
	query := `SELECT id FROM users WHERE email = $1`
	var id int
	conn.QueryRow(query, &email).Scan(&id)
	return id == 0
}

package userServices

import "lol-player-finder/database"

func DeleteUserService(id string) bool {
	conn := database.GetDb()
	query := `DELETE FROM users WHERE id = $1`
	res, _ := conn.Exec(query, id)
	lines, _ := res.RowsAffected()
	if lines == 0 {
		return false
	}
	return true
}

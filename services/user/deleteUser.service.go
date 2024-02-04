package userServices

import (
	"fmt"
	"lol-player-finder/database"
)

func DeleteUserService(id string) bool {
	conn := database.GetDb()
	query := `DELETE FROM users WHERE id = $1`
	fmt.Println(id)
	res, err := conn.Exec(query, id)
	if err != nil {
		fmt.Println(err)
	}
	lines, _ := res.RowsAffected()
	fmt.Println(lines)
	return lines > 0
}

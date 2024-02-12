package userServices

import (
	"lol-player-finder/database"
	appError "lol-player-finder/error"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func Login(email string, password string) (string, error) {
	var databaseData struct {
		Id       string `json:"id"`
		Password string `json:"password"`
	}
	query := `
	    SELECT 
		id, password
		FROM users 
		WHERE email = $1`
	conn := database.GetDb()
	conn.QueryRow(query, email).Scan(&databaseData.Id, &databaseData.Password)
	if databaseData.Password != password {
		return "", appError.BadRequestError("Invalid email or password")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    databaseData.Id,
		"email": email,
	})
	secret := os.Getenv("JWT_SECRET")
	tokenString, _ := token.SignedString([]byte(secret))

	return tokenString, nil
}

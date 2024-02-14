package middlewares

import (
	appError "lol-player-finder/error"
	"lol-player-finder/types"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func IsUserLogged(c *gin.Context) {
	token := strings.Split(c.Request.Header.Get("Authorization"), " ")[1]
	data, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			c.Error(appError.UnauthorizedError("Invalid token"))
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.Error(appError.UnauthorizedError("Invalid token"))
		return
	}
	claims := data.Claims.(jwt.MapClaims)
	var user types.UserContext
	user.Id = claims["id"].(string)
	user.Email = claims["email"].(string)
	c.Set("user", user)
	c.Next()
}

package userController

import (
	appError "lol-player-finder/error"
	userServices "lol-player-finder/services/user"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userCredentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	c.Bind(&userCredentials)
	if userCredentials.Email == "" || userCredentials.Password == "" {
		c.Error(appError.BadRequestError("Missing user credentials"))
		return
	}
	token, err := userServices.Login(userCredentials.Email, userCredentials.Password)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, token)
}

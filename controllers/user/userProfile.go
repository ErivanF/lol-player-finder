package userController

import (
	appError "lol-player-finder/error"
	userServices "lol-player-finder/services/user"
	"lol-player-finder/types"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	context, exists := c.Get("user")
	if !exists {
		c.Error(appError.NotFoundError("User not found"))
	}
	user := context.(types.UserContext)
	data := userServices.GetUserService(user.Id)

	c.JSON(200, data)
}

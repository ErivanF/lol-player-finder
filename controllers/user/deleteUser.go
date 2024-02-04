package userController

import (
	appError "lol-player-finder/error"
	userServices "lol-player-finder/services/user"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id := c.Param("id")
	deleted := userServices.DeleteUserService(id)
	if deleted {
		c.JSON(204, nil)
	} else {
		c.Error(appError.NotFoundError("user not found"))
	}
}

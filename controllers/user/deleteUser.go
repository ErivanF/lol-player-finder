package userController

import (
	userServices "lol-player-finder/services/user"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id := c.Param(":id")
	deleted := userServices.DeleteUserService(id)
	if deleted {
		c.JSON(202, nil)
	} else {
		c.JSON(404, nil)
	}
}

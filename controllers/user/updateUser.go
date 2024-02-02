package userController

import (
	userServices "lol-player-finder/services/user"
	"lol-player-finder/types"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	id := c.Param("id")

	var body types.UserChanges
	c.Bind(&body)
	user, err := userServices.UpdateUserService(id, body)
	if err != nil {
		c.Error(err)
	}
	c.JSON(200, user)
}

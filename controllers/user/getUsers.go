package userController

import (
	userServices "lol-player-finder/services/user"

	"github.com/gin-gonic/gin"
)

// @Summary Get accounts
// @Description Get all accounts
// @Produce json
// @Success 200 {object} []types.User
// @Router /user [get]
func GetAll(c *gin.Context) {
	users := userServices.GetUsersService()
	c.JSON(200, users)
}

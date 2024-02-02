package userController

import (
	userServices "lol-player-finder/services/user"

	"github.com/gin-gonic/gin"
)

// @Summary Get account
// @Description Get one account
// @Param id path string true "Your email"
// @Produce json
// @Success 200 {object} types.User
// @Router /user/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	user := userServices.GetUserService(id)
	c.JSON(200, user)
}

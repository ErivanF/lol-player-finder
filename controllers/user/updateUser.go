package userController

import (
	userServices "lol-player-finder/services/user"
	"lol-player-finder/types"

	"github.com/gin-gonic/gin"
)

// @Summary Update account
// @Description Update account
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Param name body string false "New name"
// @Param password body string false "New password"
// @Param gameName body string false "New in game name"
// @Success 200 {object} types.User
// @Failure 400 {object} appError.AppError
// @Failure 404 {object} appError.AppError
// @Router /user/{id} [patch]
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

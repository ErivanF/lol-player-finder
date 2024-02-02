package userController

import (
	appError "lol-player-finder/error"
	userServices "lol-player-finder/services/user"
	types "lol-player-finder/types"

	"github.com/gin-gonic/gin"
)

// @Summary Create account
// @Description Create a new account
// @Accept json
// @Produce json
// @Param name body string true "Your name"
// @Param email body string true "Your email"
// @Param password body string true "Your password"
// @Param gameName body string true "In game name"
// @Success 201 {object} types.User
// @Failure 400 {object} appError.AppError
// @Failure 409 {object} appError.AppError
// @Router /user [post]
func Create(c *gin.Context) {
	var body types.UserCreate
	c.Bind(&body)
	if body.Name == "" {
		c.Error(appError.BadRequestError("missing name"))
		return
	}
	if body.Email == "" {
		c.Error(appError.BadRequestError("missing email"))
		return
	}
	if !userServices.IsEmailAvailableService(body.Email) {
		c.Error(appError.ConflictError("email is not available"))
		return
	}
	if body.Password == "" {
		c.Error(appError.BadRequestError("missing password"))
		return
	}
	if body.GameName == "" {
		c.Error(appError.BadRequestError("missing gameName"))
		return
	}
	newUser, err := userServices.CreateUserService(body)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(201, newUser)
}

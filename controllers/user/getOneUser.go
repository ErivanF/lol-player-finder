package userController

import (
	"fmt"
	"lol-player-finder/error"
	userServices "lol-player-finder/services/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	fmt.Println("Entered route")
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.Error(error.BadRequestError("Invalid user id"))
		return
	}
	user := userServices.GetUserService(id)
	c.JSON(200, user)
}

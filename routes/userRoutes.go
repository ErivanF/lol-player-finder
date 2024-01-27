package routes

import (
	userController "lol-player-finder/controllers/user"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(rg *gin.Engine) {
	users := rg.Group("/user")
	users.POST("", userController.Create)
	users.GET("", userController.GetAll)
}

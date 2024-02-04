package routes

import (
	userController "lol-player-finder/controllers/user"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(rg *gin.Engine) {
	users := rg.Group("/user")
	users.POST("", userController.Create)
	users.GET("", userController.GetAll)
	users.GET("/:id", userController.GetUser)
	users.PATCH("/:id", userController.Update)
	users.DELETE("/:id", userController.Delete)
}

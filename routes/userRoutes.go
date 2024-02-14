package routes

import (
	userController "lol-player-finder/controllers/user"
	"lol-player-finder/middlewares"

	"github.com/gin-gonic/gin"
)

func AddUserRoutes(router *gin.Engine) {
	router.POST("/login", userController.Login)
	users := router.Group("/user")
	users.POST("", userController.Create)
	users.GET("", middlewares.IsUserLogged, userController.GetAll)
	users.GET("/:id", middlewares.IsUserLogged, userController.GetUser)
	users.PATCH("/:id", middlewares.IsUserLogged, userController.Update)
	users.DELETE("/:id", middlewares.IsUserLogged, userController.Delete)
}

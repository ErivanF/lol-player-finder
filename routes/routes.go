package routes

import (
	user_controller "main/controllers/user"

	"github.com/gin-gonic/gin"
)

func Start(app *gin.Engine) {
	app.POST("/user", user_controller.Create)
}

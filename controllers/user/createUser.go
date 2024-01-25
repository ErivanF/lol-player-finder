package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	c.JSON(http.StatusOK, "Response")
}

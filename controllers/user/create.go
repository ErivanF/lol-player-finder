package user_controller

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var new_user models.UserCreateRequest
	err := c.BindJSON(&new_user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	// Adicionar ao DB
	c.IndentedJSON(http.StatusOK, new_user)
}

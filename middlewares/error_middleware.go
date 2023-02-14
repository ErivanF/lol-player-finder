package middlewares

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, recovered interface{}) {
	fmt.Println(reflect.TypeOf(recovered))
	if err, ok := recovered.(string); ok {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}

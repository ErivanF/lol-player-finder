package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Helo")
	router := gin.Default()
	router.GET("/hello", hello)
	fmt.Println("Working")
	router.Run()
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"response": "Hello world"})
}

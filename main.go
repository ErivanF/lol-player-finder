package main

import (
	"fmt"
	"main/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading dotenv")
	}
	conn, err := db.OpenConnection()
	db.Migrate(conn)
	router := gin.Default()
	router.GET("/hello", hello)
	router.Run()
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"response": "Hello world"})
}

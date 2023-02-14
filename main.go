package main

import (
	"fmt"
	"main/db"
	"main/middlewares"
	"main/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading dotenv")
	}
	conn, err := db.OpenConnection()
	if err != nil {
		fmt.Println("Error connecting to database")
	}
	db.Migrate(conn)
	app := gin.Default()
	app.Use(gin.CustomRecovery(middlewares.ErrorHandler))
	routes.Start(app)
	app.Run()
}

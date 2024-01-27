package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"lol-player-finder/database"
	_ "lol-player-finder/docs"
	"lol-player-finder/middlewares"
	"lol-player-finder/routes"
)

// @Title 			Lol Player Finder
// @Version 		0.1
// @description 	API to find League of Legends Players according to their roles and championms

// @contact.name 	Erivan Ferreira
// @contact.email 	erivan.c39@gmail.com
func main() {
	godotenv.Load()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middlewares.ErrorHandler())
	database.OpenDatabase()
	database.MigrateUp()
	routes.AddUserRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":5000")
}

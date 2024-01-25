package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middlewares.ErrorHandler())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"messege": "pong",
		})
	})
	routes.AddUserRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":5000")
}

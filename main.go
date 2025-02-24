package main

import (
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	"github.com/gin-gonic/gin"
	"joaovictorliz.com/api_gocurrency/database"
	"joaovictorliz.com/api_gocurrency/routes"
	"joaovictorliz.com/api_gocurrency/services"
	_ "joaovictorliz.com/api_gocurrency/docs"
)

// @title Currency Conversion API
// @version 1.0
// @description This is an API for currency conversion with historical records.
// @host localhost:8080
// @BasePath /

func main() {
	database.InitDB()
	services.Loadenv()
	server := gin.Default()
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.MainRoutes(server)

	server.Run(":8080")
}
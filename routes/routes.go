package routes

import (
	"github.com/gin-gonic/gin"
	"joaovictorliz.com/api_gocurrency/controllers"
)

func MainRoutes(server *gin.Engine) {
	server.POST("/convert", controllers.Convert)
	server.GET("/rates/:currency", controllers.LatestCurrency)
	server.GET("/history", controllers.GetCurrencyHistory)
}
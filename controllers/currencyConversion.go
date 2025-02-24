package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"joaovictorliz.com/api_gocurrency/models"
)


// Convert godoc
// @Summary Converts a value between currencies
// @Description Converts a value from a base currency to a target currency using the current exchange rate
// @Tags Currency
// @Accept json
// @Produce json
// @Param request body models.CurrencySwagger true "Currency conversion request example"
// @Router /convert [post]
func Convert(context *gin.Context) {
	var currency models.Currency
	err := context.ShouldBindJSON(&currency)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error" : err.Error()})
		return
	}

	result, err := currency.GetCurrencyConversion()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	err = currency.SaveCurrencyHistory(result.ConvertedAmount, result.ExchangeRate)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "It was not possible save currency history."})
		return
	}

	if result == (models.Currency{}) {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Currency not found or don't exists, try again."})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"result": result})
}


// LatestCurrency godoc
// @Summary Retrieves the latest exchange rate for a specific currency
// @Description Returns the most recent exchange rate for the requested currency
// @Tags Currency
// @Produce json
// @Param currency path string true "Currency code (e.g., USD, EUR, BRL)"
// @Router /rates/{currency} [get]
func LatestCurrency(context *gin.Context) {
	currency := context.Param("currency")

	latestCurrency, err := models.GetLatestCurrency(currency)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch latest currency.", "error" : err})
		return
	}

	context.JSON(http.StatusOK, latestCurrency)
}

// GetCurrencyHistory godoc
// @Summary Returns the conversion history
// @Description Retrieves the complete history of stored currency conversions
// @Tags Currency
// @Produce json
// @Router /history [get]
func GetCurrencyHistory(context *gin.Context) {
	history, err := models.GetCurrencyHistory()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Could not fetch currency history"})
		return 
	}

	context.JSON(http.StatusOK, gin.H{"result" : history})
}
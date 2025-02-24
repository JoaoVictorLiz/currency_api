package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"joaovictorliz.com/api_gocurrency/controllers"
	"joaovictorliz.com/api_gocurrency/database"
)

type Response struct {
	Result struct {
		BaseCode         string    `json:"base_code"`
		TargetCode       string    `json:"target_code"`
		ExchangeRate     float64   `json:"conversion_rate"`
		ConvertedAmount  float64   `json:"conversion_result"`
		DateTime         time.Time `json:"date_time"`
	} `json:"result"`
}

type ResponseHist struct {
	Result []struct {
		BaseCode         string    `json:"base_code"`
		TargetCode       string    `json:"target_code"`
		ExchangeRate     float64   `json:"conversion_rate"`
		ConvertedAmount  float64   `json:"conversion_result"`
		DateTime         time.Time `json:"date_time"`
	} `json:"result"`
}



type ResponseRates struct {
	ConversionRates map[string]float64 `json:"conversion_rates"`
}

func TestGetCurrencyConversion(t *testing.T) {
	godotenv.Load()

	database.InitDB()

	server := gin.Default()
	server.POST("/convert", controllers.Convert)

	requestBody := `{"base_code": "BRL", "target_code": "EUR", "conversion_result": 250}`
	req, _ := http.NewRequest("POST", "/convert", bytes.NewBuffer([]byte(requestBody)))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var responseData Response
	err := json.Unmarshal(resp.Body.Bytes(), &responseData)
	assert.Nil(t, err) 

	assert.Equal(t, "BRL", responseData.Result.BaseCode)
	assert.Equal(t, "EUR", responseData.Result.TargetCode)
	assert.NotZero(t, responseData.Result.ExchangeRate) 
	assert.NotZero(t, responseData.Result.ConvertedAmount)
}

func TestGetCurrencyRates(t *testing.T) {
	godotenv.Load()
	database.InitDB()
	server := gin.Default()
	server.GET("/rates/:currency", controllers.LatestCurrency)

	req, _ := http.NewRequest("GET", "/rates/BRL", nil)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	var responseData ResponseRates
	err := json.Unmarshal(resp.Body.Bytes(), &responseData)
	assert.Nil(t, err) 



	expectedCurrencies := []string{"USD", "EUR", "BRL"}
	for _, currency := range expectedCurrencies {
		_, exists := responseData.ConversionRates[currency]
		assert.True(t, exists, "The currency %s should be present in the response", currency)
	}

	
	for currency, rate := range responseData.ConversionRates {
		assert.Greater(t, rate, 0.0, "The conversion rate to %s should be greater then zero", currency)
	}

}

func TestGetCurrencyHistory(t *testing.T) {
	godotenv.Load()
	database.InitDB()

	server := gin.Default()
	server.GET("/history", controllers.GetCurrencyHistory)

	req, _ := http.NewRequest("GET", "/history", nil)
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)


	var responseData ResponseHist
	err := json.Unmarshal(resp.Body.Bytes(), &responseData)
	assert.Nil(t, err)

	
	assert.NotEmpty(t, responseData.Result)

	firstHistory := responseData.Result[0] 
	assert.NotEmpty(t, firstHistory.BaseCode)
	assert.NotEmpty(t, firstHistory.TargetCode)
	assert.NotZero(t, firstHistory.ExchangeRate)
	assert.NotZero(t, firstHistory.ConvertedAmount)
}
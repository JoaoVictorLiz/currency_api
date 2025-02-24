package models

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"joaovictorliz.com/api_gocurrency/database"
	"joaovictorliz.com/api_gocurrency/services"
)

type Currency struct {
	BaseCurrency    string    `json:"base_code" binding:"required"` 
	TargetCurrency  string    `json:"target_code"  binding:"required" `
	ExchangeRate    float64   `json:"conversion_rate"`
	ConvertedAmount float64   `json:"conversion_result" binding:"required"`
	DateTime		time.Time `json:"date_time,omitempty"`
}

type CurrencySwagger struct {
	BaseCurrency    string    `json:"base_code" binding:"required" example:"BRL"` 
	TargetCurrency  string    `json:"target_code"  binding:"required" example:"EUR"`
	ConvertedAmount float64   `json:"conversion_result" binding:"required" example:"250"`
}

type LatestCurrency struct {
	ConversionRates map[string]float64 `json:"conversion_rates"`
}


func (c *Currency) GetCurrencyConversion() (Currency, error) {
	baseUrl := services.GetBaseURL()
	fullURL := baseUrl + fmt.Sprintf("/pair/%v/%v/%v", c.BaseCurrency, c.TargetCurrency, c.ConvertedAmount)
	resp, err := http.Get(fullURL)
	
	if err != nil {
		return Currency{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) 
	if err != nil {
		return Currency{}, err
	}

	var result Currency
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Currency{}, err
	}
	
	return result, nil
}

func GetLatestCurrency(currency string) (LatestCurrency, error) {
	baseUrl := services.GetBaseURL()
	fullURL := baseUrl + fmt.Sprintf("/latest/%v", currency)
	log.Println(fullURL)
	resp, err := http.Get(fullURL)
	
	if err != nil {
		return LatestCurrency{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) 
	if err != nil {
		return LatestCurrency{}, err
	}

	var result LatestCurrency
	err = json.Unmarshal(body, &result)
	if err != nil {
		return LatestCurrency{}, err
	}

	return result, nil
}

func (c *Currency) SaveCurrencyHistory(covertedResult, conversionRate float64) error {
	query := `INSERT INTO currency_hist (currencyfrom ,currencyto ,amount ,rate, result) VALUES (?, ?, ?, ?, ?)`

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.BaseCurrency, c.TargetCurrency, c.ConvertedAmount, conversionRate, covertedResult)
	return err
}

func GetCurrencyHistory() ([]Currency, error) {
	query := `SELECT currencyfrom, currencyto, rate, result, dateTime FROM currency_hist`
	rows, err := database.DB.Query(query)
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var currencies []Currency

	for rows.Next() {
		var currency Currency

		err := rows.Scan(&currency.BaseCurrency, &currency.TargetCurrency, &currency.ExchangeRate, &currency.ConvertedAmount, &currency.DateTime)
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, currency)
	}

	return currencies, nil
}
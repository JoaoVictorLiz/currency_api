package database

import (
	"database/sql"
	_"modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "currency.db")

	if err != nil {
		panic("Could not connect to database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}


func createTables() {
	createHistoryTable := `
	CREATE TABLE IF NOT EXISTS currency_hist (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		currencyfrom TEXT NOT NULL,
		currencyto TEXT NOT NULL,
		amount NUMERIC NOT NULL,
		rate NUMERIC(15,6),
		result NUMERIC NOT NULL,
		dateTime DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err := DB.Exec(createHistoryTable)

	if err != nil {
		panic("Could not create users table: " + err.Error())	
	}
}
# Currency API 🏦💱
Currency_API is a currency conversion API with transaction history, developed in Go.

## ✨ Technologies
- Go (Gin Framework)
- SQLite
- Swagger for API documentation
- Unit Tests (main_test.go)

## 📖 How to Run Locally
1. Clone the repository:
   ```sh
   git clone https://github.com/JoaoVictorLiz/currency_api.git
   cd currency_api
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Start the database and run the project:
   ```sh
   go run main.go
   ```
4. Access the Swagger documentation:
   ```
   http://localhost:8080/swagger/documentation/index.html
   ```

## 🔧 Available Endpoints
| Method | Route             | Description                    |
|--------|------------------|--------------------------------|
| `POST` | `/convert`       | Converts values between currencies |
| `GET`  | `/rates/:currency` | Retrieves the latest exchange rate |
| `GET`  | `/history`       | Fetches the conversion history |

## 👉 Example Request
### Convert Currency
**Request:**
```json
{
    "base_code": "BRL",
    "target_code": "EUR",
    "conversion_result": 250
}
```

**Response:**
```json
{
    "result": {
        "base_code": "BRL",
        "target_code": "EUR",
        "conversion_rate": 0.0063,
        "conversion_result": 1.575,
        "date_time": "2024-02-20T12:00:00Z"
    }
}
```

## 📢 Running Tests

The project includes unit tests to ensure functionality.
Run the tests with:
  ```sh
  go test ./...
   ```

## 📢 Contributing
Feel free to submit issues and pull requests!

## 📣 Contact
- LinkedIn: [My LinkedIn](https://www.linkedin.com/in/joão-victor-liz-da-silveira-b347a71b5/)
- Email: joaovictorlizsilveira@gmail.com


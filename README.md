# simple-currency-converter-api
## Getting Started

To get started with this project, follow the [Getting Started](getting_started.md) guide.

## API Documentation

### `/latest` Endpoint

- **Method**: GET
- **Description**: Retrieves the latest currency exchange rates from the API and save/update the rate in the database
- **Parameters**:
  - None
- **Response**:
  - Success: HTTP 200 OK
    ```json
    {
      "base": "EUR",
      "date": "2024-02-23",
      "rates": {
        "AED": 3.972954,
        "AFN": 79.702897,
        ...
      }
    }
    ```
  - Failure: HTTP 500 Internal Server Error

### `/convert` Endpoint

- **Method**: POST
- **Description**: Converts currency based on the provided parameters.
- **Parameters**:
  - `base` (string): Base currency code
  - `amount` (integer): Amount to convert
  - `currencies` (array of strings): Target currency codes
- **Request Example**:
  ```json
  {
    "base": "EUR",
    "amount": 100,
    "currencies": ["USD", "GBP", "JPY"]
  }
- **Response**:
 - Success: HTTP 200 OK
 ```json
 {
  "base": {
    "symbol": "EUR",
    "name": "Euro",
    "rate": 1
  },
  "rates": [
    {
      "symbol": "USD",
      "name": "US Dollar",
      "rate": 1.123456,
      "converted_amount": 112.3456
    },
    {
      "symbol": "GBP",
      "name": "British Pound",
      "rate": 0.876543,
      "converted_amount": 87.6543
    },
    ...
  ]
}
    ```
  - Failure: HTTP 500 Internal Server Error

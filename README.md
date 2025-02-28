# Password Generator API
This is a simple password generator API built with Go and the Gin framework. 
It allows users to generate random passwords with configurable options, including letters, special characters, and numbers.

## Features
* Generate random passwords with specified length
* Include or exclude letters, special characters, and numbers
* Ensures password validity based on provided criteria
* Configurable API port via environment variables

## Requirements
* Go 1.18+
* Gin framework

## Installation
Clone the repository and navigate to the project directory:
```
git clone https://github.com/ksl28/password-generator.git
cd password-generator
```

Install dependencies:
```
go mod tidy
```

## Usage
Run the API:

```sh
# For Linux
export apiport=8080

# For PowerShell
$env:apiport="8080"

# Run the application
go run main.go
```

## API Endpoint
Generate a Password:
```sh
# Default
GET /api/v1/genpwd

# With special settings
GET /api/v1/genpwd?length=16&includeLetter=true&includeSpecial=true&includeNumbers=true
```

## Example output
```
w^f!4u2S
```


## Query Parameters
| Parameter       | Type  | Default | Required | Description                              |
|---------------|------|---------|----------|----------------------------------|
| `length`       | int  | 16      | No       | Length of the password (minimum: 8) |
| `includeLetter` | bool | true    | No       | Include letters in the password   |
| `includeSpecial` | bool | true    | No       | Include special characters in the password |
| `includeNumbers` | bool | true    | No       | Include numbers in the password  |

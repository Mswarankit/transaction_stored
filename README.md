# Transaction_stored


This project implements a RESTful web service for managing financial transactions. It allows storing transactions with types and amounts, linking transactions, and retrieving transaction information.

#### Features

- Create and store transactions with amount, type, and optional parent ID
- Retrieve transaction details by ID
- Get all transaction IDs of a specific type
- Calculate the sum of all transactions linked to a particular transaction

#### Prerequisites
- Go (version 1.22 or later)
- Git

###### Steps
- clone the repo `https://github.com/Mswarankit/transaction_stored`
- `cd transaction-stored`

###### Install dependencies
- `go mod tidy`

###### Build the project
- `go build -o transaction-service cmd/main.go`

###### Run the server
- `./transaction-service`

###### Testing
- `go test`


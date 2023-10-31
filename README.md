# Bankie.go

Bankie.go is a simple banking application built with Go programming language. It provides basic banking functionalities such as creating accounts, making deposits and withdrawals, and transferring funds between accounts.

## Features

- **Account Management**: Users can create new bank accounts and manage their existing accounts.
- **Deposits and Withdrawals**: Users can deposit money into their accounts or make withdrawals.
- **Fund Transfers**: Users can transfer funds from one account to another.
- **Transaction History**: The application maintains a transaction history for each account, allowing users to track their transactions.

## Setup with Docker Compose

To run Bankie.go using Docker Compose, make sure you have Docker and Docker Compose installed on your system. Then, follow these steps:

1. Clone the repository:

   ```shell
   git clone https://github.com/theghostmac/bankie.go.git
   ```

2. Navigate to the project directory:

   ```shell
   cd bankie.go
   ```

3. Start the application using Docker Compose:

   ```shell
   docker-compose up
   ```

This will start the application and its dependencies, including PostgreSQL, using Docker Compose.

## Example Usage with cURL

Once the application is running, you can interact with it using a RESTful API. Here are some example cURL commands:

### Create a new bank account:

```shell
curl -X POST -H "Content-Type: application/json" -d '{
    "FirstName": "John",
    "LastName": "Doe",
    "Email": "johndoe@example.com"
}' http://localhost:8082/accounts
```

Output:

```json
{"firstName":"John","lastName":"Doe","email":"johndoe@example.com"}
```

### Retrieve information about a specific account:

Replace `:id` with the actual account ID:

```shell
curl http://localhost:8082/accounts/:id
```

### Other Endpoints:

- **POST /accounts/:id/deposit**: Make a deposit to a specific account. Provide the deposit amount.
- **POST /accounts/:id/withdraw**: Make a withdrawal from a specific account. Provide the withdrawal amount.
- **POST /accounts/:id/transfer**: Transfer funds from one account to another. Provide the recipient account ID and the transfer amount.
- **GET /accounts/:id/transactions**: Retrieve the transaction history of a specific account.

Make sure to replace `:id` with the actual account ID in the endpoint URLs.

## Contributing

Contributions to Bankie.go are welcome! If you find any issues or would like to add new features, please feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](#).
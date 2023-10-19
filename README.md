# Stori Transaction Summary

## Project description

For this challenge you must create a system that processes a file from a mounted directory. The file
will contain a list of debit and credit transactions on an account. Your function should process the file
and send summary information to a user in the form of an email.

Credit transactions are indicated with a plus sign like +60.5. Debit transactions are indicated by a minus sign like -20.46
We prefer that you code in Python or Golang; but other languages are ok too. Package your code in one or more Docker images. Include any build or run scripts, Dockerfiles or docker-compose files needed to build and execute your code.

Bonus points

1. Save transaction and account info to a database
2. Style the email and include Stori’s logo
3. Package and run code on a cloud platform like AWS. Use AWS Lambda and S3 in lieu of Docker.

Delivery and code requirements

Your project must meet these requirements:

1. The summary email contains information on the total balance in the account, the number of transactions grouped by month, and the average credit and average debit amounts grouped by month. Using the transactions in the image above as an example, the summary info would be
   Total balance is 39.74
   Number of transactions in July: 2
   Number of transactions in August: 2
   Average debit amount: -15.38
   Average credit amount: 35.25

2. Include the file you create in CSV format.

3. Code is versioned in a git repository. The REA

## Setup

### Prerequisites

Before proceeding, ensure that you have the following prerequisites:

- You need to have installed [docker](https://docs.docker.com/engine/install/)
- Optionally you can install [make](https://makefiletutorial.com/) to make the next steps easier

## Steps

- Run the command `make docker-build` or `docker build -t stori-transaction-summary .`
- Run the command `make docker-compose` or `docker-compose up -d`

# Golang Clean architecture

The solution architecture we can divide the code in 6 main layers:

- Models: Is a set of data structures.
- Services: Contains application specific business rules. It encapsulates and implements all the use cases of the system.
- Controllers: Is a set of adapters that convert data from the format most convenient for the use cases and models.
- Repository: Contains the database operations such as querying, inserting, updating, and deleting data; separating the data access logic from the rest of the application's business logic.
- Utils and Config: Is generally composed of frameworks and tools.
- Email: Provides an implementation for generating and sending the email with SMTP

In Clean Architecture, each layer of the application (use case, data service and domain model) only depends on interface of other layers instead of concrete types.
Dependency Injection is one of the SOLID principles, a rule about the constraint between modules that abstraction should not depend on details.
Clean Architecture uses this rule to keep the dependency direction from outside to inside.

## Routes

### Sign-Up Account

- Method: POST
- URL: `/api/sign-up`
- Description: Allows users to sign up by providing their email address. This initiates the registration process.
- Request:
  - Headers:
    - Content-Type: application/json (Required) - Specifies that the request body contains JSON data.
    - Body: JSON Object with an "email" field (Required) - The email address of the user signing up.
- Response:
  - Status Codes:
    - 200 OK: Sign-up successful.
    - 400 Bad Request: Invalid request or missing email field.
    - 500 Internal Server Error: An error occurred during sign-up process.
  - Body: JSON object containing the account information.

#### Example

```bash
curl --location 'localhost:8080/api/sign-up' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@gmail.com"
}'
```

### Upload Transactions

- Method: POST
- URL: `/api/account/:id/upload-transactions`
- Description: Allows to upload a .CSV file containing transaction data. After the file is uploaded, the server processes it to generate a financial summary report, which is then sent via email to the user and saved into the database.
- Request:
  - Headers:
    - Content-Type: multipart/form-data
  - Path Variables:
    - id: The account ID for the user we want to send the email.
  - Body: Form data with a file field named "file" - The uploaded .CSV file.
- Response:
  - Status Code:
    - 200 OK - Successful uploaded and email sent.
    - 400 Bad Request - Invalid request or uploaded file.
    - 500 Internal Server Error - An error occurred during processing.
  - Body: JSON object containing the summary of transactions.

#### Example

```bash
curl --location 'localhost:8080/api/account/1c123230-5c31-4f39-836d-fe426bbb4d2a/upload-transactions' \
--form 'file=@"///wsl.localhost/Ubuntu/root/stori_summary_transaction/tsx.csv"'
```

### Retrieve Account Transactions

- Method: GET
- URL: `localhost:8080/api/account/:id/transactions`
- Description: Retrieves the transactions for the specified account.
- Request:
  - Path Variables:
    - id: The account ID for the user we want to send the email.
- Response:
  - Status Codes:
    - 200 OK: Successful request.
    - 404 Not Found: Account or transactions not found.
    - 500 Internal Server Error - An error occurred during processing.
  - Body: JSON object containing the list of transactions.

#### Example

```bash
curl --location 'localhost:8080/api/account/1c123230-5c31-4f39-836d-fe426bbb4d2a/transactions'
```

## Built With

- [Go](https://go.dev/) - version 1.20.2
- [Testify](https://github.com/stretchr/testify)
- [Viper](https://github.com/spf13/viper)
- [Makefile](https://www.gnu.org/software/make/manual/make.html#Introduction)
- [UberFx](https://github.com/uber-go/fx)
- [Gin](https://github.com/gin-gonic-gin)
- [Logrus](https://github.com/sirupsen/logrus)

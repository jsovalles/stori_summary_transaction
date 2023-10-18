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

### Local test

## Steps
* Download all dependencies and run `go mod tidy`
* Run the command `make test` or `go clean -testcache && go test -v -cover ./...`
* Run the command `make run` or `go run cmd/main.go`

### Docker

#### Prerequisites

Before proceeding, ensure that you have the following prerequisites:

- You need to have installed [docker](https://docs.docker.com/engine/install/)

## Steps
* Run the command `make docker-build` or `docker build -t stori-transaction-summary .`
* Run the command `make docker-compose` or `docker-compose up -d`

# Golang Clean architecture

The solution architecture we can divide the code in 5 main layers:

- Models: Is a set of data structures.
- Services: Contains application specific business rules. It encapsulates and implements all the use cases of the system.
- Controllers: Is a set of adapters that convert data from the format most convenient for the use cases and models.
- Utils and Config: Is generally composed of frameworks and tools.
- Email: Provides an implementation for generating and sending the email with SMTP

In Clean Architecture, each layer of the application (use case, data service and domain model) only depends on interface of other layers instead of concrete types. 
Dependency Injection is one of the SOLID principles, a rule about the constraint between modules that abstraction should not depend on details. 
Clean Architecture uses this rule to keep the dependency direction from outside to inside.

## Routes

### upload Transactions

- Method: POST
- URL: `/api/upload-transactions`
- Description: Allows to upload a .CSV file containing transaction data. After the file is uploaded, the server processes it to generate a financial summary report, which is then sent via email to the user.
- Request:
    - Headers:
        - Content-Type: multipart/form-data
    - Body: Form data with a file field named "file".
- Response:
    - Status Code: 
        - 200 OK - Successful uploaded and email sent.
        - 400 Bad Request - Invalid request or uploaded file.
        - 500 Internal Server Error - An error occurred during processing.
    - Body: JSON object containing the summary of transactions.

#### Example

```bash
curl --location 'localhost:8080/api/upload-transactions' \
--form 'file=@"///wsl.localhost/Ubuntu/home/jsovalles/StoriTransactionSummary/tsx.csv"'
```

## Built With

- [Go](https://go.dev/) - version 1.20.2
- [Testify](https://github.com/stretchr/testify)
- [Viper](https://github.com/spf13/viper)
- [Makefile](https://www.gnu.org/software/make/manual/make.html#Introduction)
- [UberFx](https://github.com/uber-go/fx)
- [Gin](https://github.com/gin-gonic-gin)
- [Logrus](https://github.com/sirupsen/logrus)

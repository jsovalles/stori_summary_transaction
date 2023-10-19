package models

import "github.com/google/uuid"

type NewAccount struct {
	Email string `json:"email"`
}

type Account struct {
	Id    uuid.UUID `json:"id" db:"account_id"`
	Email string    `json:"email" db:"account_email"`
}
type Transaction struct {
	TransactionId int       `json:"transactionId" db:"transaction_id"`
	AccountId     uuid.UUID `json:"-" db:"account_id"`
	Month         string    `json:",omitempty"`
	Amount        float64   `json:"amount" db:"amount"`
	MonthDate     string    `json:"monthDate" db:"transaction_date"`
}

type TransactionSummary struct {
	TotalBalance float64                `json:"totalBalance"`
	MonthStats   map[string]*MonthStats `json:"monthStats"`
}

type MonthStats struct {
	TotalCredit      float64 `json:"totalCredit"`
	TotalDebit       float64 `json:"totalDebit"`
	AverageCredit    float64 `json:"averageCredit"`
	AverageDebit     float64 `json:"averageDebit"`
	TransactionCount int     `json:"transactionCount"`
}

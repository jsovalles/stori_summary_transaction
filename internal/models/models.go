package models

type Record struct {
	ID          int
	Month       string
	Transaction float64
}

type TransactionSummary struct {
	TotalBalance float64
	MonthStats   map[string]*MonthStats
}

type MonthStats struct {
	TotalCredit      float64
	TotalDebit       float64
	AverageCredit    float64
	AverageDebit     float64
	TransactionCount int
}

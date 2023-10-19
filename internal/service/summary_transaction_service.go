package service

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jsovalles/stori_transaction_summary/internal/mail"
	"github.com/jsovalles/stori_transaction_summary/internal/models"
	"github.com/jsovalles/stori_transaction_summary/internal/repository"
	"github.com/jsovalles/stori_transaction_summary/internal/utils"
	"go.uber.org/fx"
)

type SummaryTransactionService interface {
	SignUpAccount(account models.Account) (createdAccount models.Account, err error)
	SummaryTransaction(file multipart.File, accountId string) (ts models.TransactionSummary, err error)
	ListAccountTransactionsByAccountId(accountId string) (transactions []models.Transaction, err error)
}

type summaryTransactionService struct {
	email      mail.Email
	repository repository.SummaryTransactionRepository
}

func (st *summaryTransactionService) SignUpAccount(account models.Account) (createdAccount models.Account, err error) {

	createdAccount, err = st.repository.SignUpAccount(account)
	if err != nil {
		return
	}

	return
}

func (st *summaryTransactionService) SummaryTransaction(file multipart.File, accountId string) (ts models.TransactionSummary, err error) {

	parsedAccountId, err := uuid.Parse(accountId)
	if err != nil {
		return models.TransactionSummary{}, fmt.Errorf(utils.ParseUUIDErr, err)
	}

	transactions, err := fileToTransaction(file, parsedAccountId)
	if err != nil {
		return models.TransactionSummary{}, err
	}

	ts = calculateSummaryInformation(transactions)

	err = st.repository.SaveAccountTransactions(transactions)
	if err != nil {
		return models.TransactionSummary{}, err
	}

	account, err := st.repository.GetAccountByAccountId(accountId)
	if err != nil {
		return
	}

	err = st.email.SendEmail(ts, account)
	if err != nil {
		return models.TransactionSummary{}, err
	}

	return
}

func (st *summaryTransactionService) ListAccountTransactionsByAccountId(accountId string) (transactions []models.Transaction, err error) {

	transactions, err = st.repository.ListAccountTransactionsByAccountId(accountId)
	if err != nil {
		return
	}

	return
}

func fileToTransaction(file multipart.File, accountId uuid.UUID) (records []models.Transaction, err error) {

	csvRecords, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return
	}

	for i, record := range csvRecords {

		if i != 0 {

			transactionId, err := strconv.Atoi(record[0])
			if err != nil {
				return []models.Transaction{}, fmt.Errorf(utils.CsvIdErr, err)
			}

			parsedDate, err := time.Parse("1/2", record[1])
			if err != nil {
				return []models.Transaction{}, fmt.Errorf(utils.CsvDateErr, err)
			}

			month := fmt.Sprintf("%s - %s", parsedDate.Format("01"), parsedDate.Format("January"))

			amount, err := strconv.ParseFloat(record[2], 64)
			if err != nil {
				return []models.Transaction{}, fmt.Errorf(utils.CsvTransactionErr, err)
			}

			records = append(records, models.Transaction{
				TransactionId: transactionId,
				AccountId:     accountId,
				Month:         month,
				MonthDate:     record[1],
				Amount:        amount,
			})
		}
	}

	return
}

func calculateSummaryInformation(records []models.Transaction) (total models.TransactionSummary) {

	monthStatsMap := make(map[string]*models.MonthStats)
	for _, record := range records {

		monthStats, exists := monthStatsMap[record.Month]
		if !exists {
			monthStats = &models.MonthStats{}
			monthStatsMap[record.Month] = monthStats
		}

		if record.Amount > 0 {
			monthStats.TotalCredit += record.Amount
		} else {
			monthStats.TotalDebit += record.Amount
		}

		monthStats.TransactionCount++
		total.TotalBalance += record.Amount
	}

	for month, monthStats := range monthStatsMap {
		monthStatsMap[month].AverageCredit = monthStats.TotalCredit / float64(monthStats.TransactionCount)
		monthStatsMap[month].AverageDebit = monthStats.TotalDebit / float64(monthStats.TransactionCount)
	}

	total.MonthStats = monthStatsMap

	return
}

func NewSummaryTransactionService(email mail.Email, repository repository.SummaryTransactionRepository) SummaryTransactionService {
	return &summaryTransactionService{email: email, repository: repository}
}

var Module = fx.Provide(NewSummaryTransactionService)

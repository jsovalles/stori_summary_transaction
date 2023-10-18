package service

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/jsovalles/stori_transaction_summary/internal/mail"
	"github.com/jsovalles/stori_transaction_summary/internal/models"
	"github.com/jsovalles/stori_transaction_summary/internal/repository"
	"github.com/jsovalles/stori_transaction_summary/internal/utils"
	"go.uber.org/fx"
)

type SummaryTransactionService interface {
	SignUpAccount(account models.Account) (res models.Account, err error)
	SummaryTransaction(file multipart.File, accountId string) (ts models.TransactionSummary, err error)
}

type summaryTransactionService struct {
	email      mail.Email
	repository repository.SummaryTransactionRepository
}

func (st *summaryTransactionService) SignUpAccount(account models.Account) (res models.Account, err error) {

	err = st.repository.SignUpAccount(account)
	if err != nil {
		return
	}

	res = models.Account{Id: account.Id, Email: account.Email}

	return
}

func (st *summaryTransactionService) SummaryTransaction(file multipart.File, accountId string) (ts models.TransactionSummary, err error) {

	records, err := fileToRecord(file)
	if err != nil {
		return models.TransactionSummary{}, err
	}

	ts = calculateSummaryInformation(records)

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

func fileToRecord(file multipart.File) (records []models.Record, err error) {

	csvRecords, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return
	}

	for i, record := range csvRecords {

		if i != 0 {

			id, err := strconv.Atoi(record[0])
			if err != nil {
				return []models.Record{}, fmt.Errorf(utils.CsvIdErr, err)
			}

			date, err := time.Parse("1/2", record[1])
			if err != nil {
				return []models.Record{}, fmt.Errorf(utils.CsvDateErr, err)
			}
			month := date.Format("January")

			transaction, err := strconv.ParseFloat(record[2], 64)
			if err != nil {
				return []models.Record{}, fmt.Errorf(utils.CsvTransactionErr, err)
			}

			records = append(records, models.Record{
				ID:          id,
				Month:       month,
				Transaction: transaction,
			})
		}
	}

	return
}

func calculateSummaryInformation(records []models.Record) (total models.TransactionSummary) {

	monthStatsMap := make(map[string]*models.MonthStats)
	for _, record := range records {

		monthStats, exists := monthStatsMap[record.Month]
		if !exists {
			monthStats = &models.MonthStats{}
			monthStatsMap[record.Month] = monthStats
		}

		if record.Transaction > 0 {
			monthStats.TotalCredit += record.Transaction
		} else {
			monthStats.TotalDebit += record.Transaction
		}

		monthStats.TransactionCount++
		total.TotalBalance += record.Transaction
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

package repository

import (
	"database/sql"
	"fmt"

	"github.com/jsovalles/stori_transaction_summary/internal/config"
	"github.com/jsovalles/stori_transaction_summary/internal/models"
	"github.com/jsovalles/stori_transaction_summary/internal/utils"
	"go.uber.org/fx"
)

const (
	accountsTable              = "accounts"
	transactionsTable          = "transactions"
	createAccount              = "INSERT INTO " + accountsTable + " (account_email) VALUES(:account_email)"
	getAccountByAccountId      = "SELECT * FROM " + accountsTable + " WHERE account_id = $1"
	getAccountByAccountEmail   = "SELECT * FROM " + accountsTable + " WHERE account_email = $1"
	saveTransactions           = "INSERT INTO " + transactionsTable + " (transaction_id, account_id, amount, transaction_date) VALUES(:transaction_id,:account_id,:amount,:transaction_date)"
	listTransactionsByAccountId = "SELECT * FROM " + transactionsTable + " WHERE account_ID = $1"
)

type SummaryTransactionRepository interface {
	SignUpAccount(account models.Account) (createdAccount models.Account, err error)
	GetAccountByAccountId(accountId string) (account models.Account, err error)
	SaveAccountTransactions(transactions []models.Transaction) (err error)
	ListAccountTransactionsByAccountId(accountId string) (transactions []models.Transaction, err error)
}

type summaryTransactionRepository struct {
	env config.Config
	db  config.Database
}

func (r *summaryTransactionRepository) SignUpAccount(account models.Account) (createdAccount models.Account, err error) {
	_, err = r.db.Client.NamedExec(createAccount, account)
	if err != nil {
		return
	}

	err = r.db.Client.Get(&createdAccount, getAccountByAccountEmail, account.Email)
	if err != nil {
		return
	}

	return
}

func (r *summaryTransactionRepository) GetAccountByAccountId(accountId string) (account models.Account, err error) {
	err = r.db.Client.Get(&account, getAccountByAccountId, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf(utils.NoUserErr)
			return
		}
		return
	}
	return
}

func (r *summaryTransactionRepository) SaveAccountTransactions(transactions []models.Transaction) (err error) {
	_, err = r.db.Client.NamedExec(saveTransactions, transactions)
	if err != nil {
		return
	}

	return
}

func (r *summaryTransactionRepository) ListAccountTransactionsByAccountId(accountId string) (transactions []models.Transaction, err error) {
	err = r.db.Client.Select(&transactions, listTransactionsByAccountId, accountId)
	if err != nil {
		return
	}

	return
}

func NewSummaryTransactionRepository(env config.Config, db config.Database) SummaryTransactionRepository {
	return &summaryTransactionRepository{env: env, db: db}
}

var Module = fx.Provide(NewSummaryTransactionRepository)

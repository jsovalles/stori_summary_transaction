package repository

import (
	"database/sql"
	"fmt"

	"github.com/jsovalles/stori_transaction_summary/internal/config"
	"github.com/jsovalles/stori_transaction_summary/internal/models"
	"go.uber.org/fx"
)

const (
	accountsTable         = "accounts"
	createAccount         = "INSERT INTO " + accountsTable + " (account_email) VALUES(:account_email)"
	getAccountByAccountId = "SELECT * FROM " + accountsTable + " WHERE account_id = $1"
)

type SummaryTransactionRepository interface {
	SignUpAccount(account models.Account) (err error)
	GetAccountByAccountId(accountId string) (account models.Account, err error)
}

type summaryTransactionRepository struct {
	env config.Config
	db  config.Database
}

func (r *summaryTransactionRepository) SignUpAccount(account models.Account) (err error) {
	_, err = r.db.Client.NamedExec(createAccount, account)
	if err != nil {
		return
	}
	return
}

func (r *summaryTransactionRepository) GetAccountByAccountId(accountId string) (account models.Account, err error) {
	err = r.db.Client.Get(&account, getAccountByAccountId, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("there are no results for this user, please validate")
			return
		}
		return
	}
	return
}

func NewSummaryTransactionRepository(env config.Config, db config.Database) SummaryTransactionRepository {
	return &summaryTransactionRepository{env: env, db: db}
}

var Module = fx.Provide(NewSummaryTransactionRepository)

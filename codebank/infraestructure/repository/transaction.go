package repository

import (
	"database/sql"

	"github.com/codeedu/codebank/domain"
)

type TransactionRepositoryDb struct {
	Db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{Db: db}
}

func (t *TransactionRepositoryDb) SaveTransaction(transaction domain.Transaction, creditCard domain.CreditCard) error {

}

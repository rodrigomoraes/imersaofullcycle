package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionRepository interface {
	SaveTransaction(transaction Transaction, creditCard CreditCard) error
	GetCreditCard(CreditCard CreditCard) (CreditCard, error)
	CreateCreditCard(CreditCard CreditCard) error
}

type Transaction struct {
	ID           string
	Amount       float64
	Status       string
	Description  string
	Store        string
	CreditCardID string
	CreatedAt    time.Time
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	t.ID = uuid.NewV4().String()
	t.CreatedAt = time.Now()
	return t
}

func (t *Transaction) ProcessAndValidate(creditCard *CreditCard) {
	if t.Amount+creditCard.Balance > creditCard.Limit {
		t.Status = "rejected"
		return
	} else {
		t.Status = "approved"
		creditCard.Balance = creditCard.Balance + t.Amount
	}
}

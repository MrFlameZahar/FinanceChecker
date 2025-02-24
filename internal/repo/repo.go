package repo

import (
	"FinanceChecker/internal/models/transaction"
)

type Repository interface {
	Add(transaction transaction.Transaction, userID int64) (int64, error)
	Get(userID int64, transactionType string) ([]transaction.Transaction, error)
	Delete(transactionID int64) error
}

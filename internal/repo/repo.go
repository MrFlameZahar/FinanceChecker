package repo

import (
	"FinanceChecker/internal/models/transaction"
	"time"
)

type Repository interface {
	Add(transaction transaction.Transaction, userID int64) (int64, error)
	Get(dateFrom, dateTo time.Time, userID int64, transactionType string) ([]transaction.Transaction, error)
	Delete(transactionID int64) error
}

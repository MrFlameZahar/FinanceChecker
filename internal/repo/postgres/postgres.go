package postgres

import (
	"FinanceChecker/internal/models/transaction"
	"database/sql"
	"fmt"
	"log/slog"
	"time"
)

const (
	income  = "Income"
	expense = "Expense"
)

type Storage struct {
	DB *sql.DB
}

func New(log slog.Logger) *Storage {
	var db *sql.DB

	connStr := "user=mrflame password=Zaxaro12 dbname=test host=127.0.0.1 port=5432 sslmode=disable"
	var err error

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Error("cant connect to database")
		return nil
	}

	return &Storage{DB: db}
}

func (s *Storage) Add(transaction transaction.Transaction, userID int64) (int64, error) {
	switch transaction.Type {
	case income:

	case expense:

	default:
		return 0, fmt.Errorf("wrong transaction type")
	}
	return 0, nil
}

func (s *Storage) Get(dateFrom, dateTo time.Time, userID int64, transactionType string) ([]transaction.Transaction, error) {
	return nil, nil
}

func (s *Storage) Delete(transactionID int64) error {
	return nil
}

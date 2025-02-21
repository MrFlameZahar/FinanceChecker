package sqlite

import (
	"FinanceChecker/internal/models/transaction"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS user(
			id INTEGER PRIMARY KEY,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL);
	`)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	stmt, err = db.Prepare(`
	CREATE TABLE IF NOT EXISTS transactions(
		id INTEGER PRIMARY KEY,
		user id NOT NULL UNIQUE,
		amount INTEGER NOT NULL,
		date INTEGER NOT NULL,
		comment TEXT,
		type TEXT NOT NULL,
		currency TEXT NOT NULL);
	`)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Add(transaction transaction.Transaction, userID int64) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO transactions(id, user, amount, date, comment, type, currency) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}
	timestump := transaction.Date.Unix()
	result, err := stmt.Exec(transaction.ID, userID, transaction.Amount, timestump, transaction.Comment, transaction.Type, transaction.Currency)
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}

	return id, nil
}
func (s *Storage) Get(dateFrom, dateTo time.Time, userID int64, transactionType string) ([]transaction.Transaction, error) {
	panic("implement this")
}
func (s *Storage) Delete(transactionID int64) error {
	panic("implement this")
}

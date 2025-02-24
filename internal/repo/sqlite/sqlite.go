package sqlite

import (
	"FinanceChecker/internal/models/transaction"
	"database/sql"
	"fmt"

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
		user_id NOT NULL,
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
	stmt, err := s.db.Prepare("INSERT INTO transactions(user_id, amount, date, comment, type, currency) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}
	result, err := stmt.Exec(userID, transaction.Amount, transaction.Date, transaction.Comment, transaction.Type, transaction.Currency)
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}

	return id, nil
}
func (s *Storage) Get(userID int64, transactionType string) ([]transaction.Transaction, error) {
	stmt, err := s.db.Prepare("SELECT * FROM transactions WHERE user_id = ? AND type = ?")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	var transactions []transaction.Transaction

	rows, err := stmt.Query(userID, transactionType)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var transaction transaction.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.Amount, &transaction.Date, &transaction.Comment, &transaction.Type, &transaction.Currency); err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		transactions = append(transactions, transaction)
	}

	return transactions, err
}
func (s *Storage) Delete(transactionID int64) error {
	panic("implement this")
}

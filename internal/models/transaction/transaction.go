package transaction

import "time"

type Transaction struct {
	Amount   int
	Date     time.Time
	Comment  string
	Currency string
	Type     string
	ID       int64
}

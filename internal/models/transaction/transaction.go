package transaction

type Transaction struct {
	Amount   int
	Date     int64
	Comment  string
	Currency string
	Type     string
	ID       int64
	UserID   int64
}

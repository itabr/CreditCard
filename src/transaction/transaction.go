package transaction

type Transaction struct {
	amount float32
	date   int
}

// NewTransaction creates new Transaction which holds the values amount(float32) and date(int)
func NewTransaction(Amount float32, Date int) *Transaction {
	var t = Transaction{
		amount: Amount,
		date:   Date}
	return &t
}

// GetAmount returns the value amount(float32)
func (t *Transaction) GetAmount() float32 {
	return t.amount
}

// GetDate returns the value date(int)
func (t *Transaction) GetDate() int {
	return t.date
}

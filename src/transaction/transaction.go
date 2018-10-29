package transaction

type Transaction struct {
	amount float32
	date   int
}

func NewTransaction(Amount float32, Date int) *Transaction {
	var t = Transaction{
		amount: Amount,
		date:   Date}
	return &t
}

func (t *Transaction) GetAmount() float32 {
	return t.amount
}

func (t *Transaction) GetDate() int {
	return t.date
}

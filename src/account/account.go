package account

type Account struct {
	id                 int
	outstandingBalance float32
	interest           float32
	apr                float32
	creditLimit        float32
	cardSwipes         []float32
	payments           []float32
}

func (a *Account) Getid() int {
	return a.id
}

func (a *Account) GetoutstandingBalance() float32 {
	return a.outstandingBalance
}

func (a *Account) Getinterest() float32 {
	return a.interest
}

func (a *Account) Getapr() float32 {
	return a.apr
}

func (a *Account) GetcreditLimit() float32 {
	return a.creditLimit
}

func (a *Account) GetcardSwipes() []float32 {
	return a.cardSwipes
}

func (a *Account) Getpayments() []float32 {
	return a.payments
}

func (a *Account) Charge(c float32) {
	a.cardSwipes = append(a.cardSwipes, c)
	a.outstandingBalance = a.outstandingBalance + c
}

func (a *Account) Pay(c float32) {
	a.payments = append(a.payments, c)
	a.outstandingBalance = a.outstandingBalance - c
}

func NewAccount(ID int, Apr float32, CreditLimit float32) *Account {
	var a = Account{
		id:          ID,
		apr:         Apr,
		creditLimit: CreditLimit}
	return &a
}

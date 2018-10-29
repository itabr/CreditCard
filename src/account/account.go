package account

import (
	"../transaction"
)

// Account struct represent credit card account
type Account struct {
	id                 int
	outstandingBalance float32
	balance            float32
	interest           float32
	apr                float32
	creditLimit        float32
	cardSwipes         []*transaction.Transaction
	payments           []*transaction.Transaction
	dateOpen           int
}

// Getid returns the value id (int)
func (a *Account) Getid() int {
	return a.id
}

// GetoutstandingBalance returns the value outstandingBalance (float32)
func (a *Account) GetoutstandingBalance() float32 {
	return a.outstandingBalance
}

// Getbalance returns the value balance (float32)
func (a *Account) Getbalance() float32 {
	return a.balance
}

// Getinterest returns the value interest(float32)
func (a *Account) Getinterest() float32 {
	return a.interest
}

// Getapr returns the value apr(float32)
func (a *Account) Getapr() float32 {
	return a.apr
}

// GetcreditLimit returns the value creditLimit(float32)
func (a *Account) GetcreditLimit() float32 {
	return a.creditLimit
}

// GetcardSwipes returns the value cardSwipes([]*transaction.Transaction)
func (a *Account) GetcardSwipes() []*transaction.Transaction {
	return a.cardSwipes
}

// Getpayments returns the value payments([]*transaction.Transaction)
func (a *Account) Getpayments() []*transaction.Transaction {
	return a.payments
}

// GetdateOpen returns the value dateOpen(int)
func (a *Account) GetdateOpen() int {
	return a.dateOpen
}

// NewAccount creates new account
func NewAccount(id int, apr float32, creditLimit float32, dateOpen int) *Account {
	var a = Account{
		id:          id,
		apr:         apr,
		creditLimit: creditLimit,
		dateOpen:    dateOpen}
	return &a
}

func (a *Account) updateBalance(c float32) {
	a.balance = a.balance + c
	a.outstandingBalance = a.outstandingBalance + c
}

// Charge apply a Charge with value c to account at date d
func (a *Account) Charge(c float32, d int) {
	a.cardSwipes = append(a.cardSwipes, transaction.NewTransaction(c, d))
	a.updateBalance(c)
}

// Payment apply a Payment with value c to account at date d
func (a *Account) Payment(c float32, d int) {
	a.payments = append(a.payments, transaction.NewTransaction(c, d))
	a.updateBalance(-1 * c)
}

// ApplyDailyInterest apply a new daily into interest
func (a *Account) ApplyDailyInterest() {
	a.interest = a.interest + (a.balance*a.apr)/365.0
}

// UpdateOutstandingBalance apply a monthly interest into outstanding balance
func (a *Account) UpdateOutstandingBalance() {
	a.outstandingBalance = a.interest + a.outstandingBalance
	a.interest = 0
}

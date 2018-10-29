package account

import (
	"../transaction"
)

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

func (a *Account) Getid() int {
	return a.id
}

func (a *Account) GetoutstandingBalance() float32 {
	return a.outstandingBalance
}

func (a *Account) Getbalance() float32 {
	return a.balance
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

func (a *Account) GetcardSwipes() []*transaction.Transaction {
	return a.cardSwipes
}

func (a *Account) Getpayments() []*transaction.Transaction {
	return a.payments
}

func (a *Account) GetdateOpen() int {
	return a.dateOpen
}

func NewAccount(id int, apr float32, creditLimit float32, dateOpen int) *Account {
	var a = Account{
		id:          id,
		apr:         apr,
		creditLimit: creditLimit,
		dateOpen:    dateOpen}
	return &a
}

func (a *Account) UpdateBalance(c float32) {
	a.balance = a.balance + c
	a.outstandingBalance = a.outstandingBalance + c
}

func (a *Account) Charge(c float32, d int) {
	a.cardSwipes = append(a.cardSwipes, transaction.NewTransaction(c, d))
	a.UpdateBalance(c)
}

func (a *Account) Payment(c float32, d int) {
	a.payments = append(a.payments, transaction.NewTransaction(c, d))
	a.UpdateBalance(-1 * c)
}

func (a *Account) ApplyDailyInterest() {
	a.interest = a.interest + (a.balance*a.apr)/365.0
}

func (a *Account) UpdateOutstandingBalance() {
	a.outstandingBalance = a.interest + a.outstandingBalance
	a.interest = 0
}

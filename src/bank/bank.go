package bank

import (
	"fmt"

	"../account"
)

type Bank struct {
	accountList []*account.Account
	date        int
}

func (b *Bank) CreateAccount(apr float32, creditLimit float32) {
	var id = len(b.accountList)
	var c = account.NewAccount(id, apr, creditLimit, b.date)
	b.accountList = append(b.accountList, c)
}

func (b *Bank) findAccount(id int) *account.Account {
	return b.accountList[id]
}

func (b *Bank) Charge(id int, c float32) {
	b.findAccount(id).Charge(c, b.date)
}

func (b *Bank) MakePayment(id int, c float32) {
	b.findAccount(id).Payment(c, b.date)
}

func (b *Bank) PrintAccountInfo(id int) {
	fmt.Printf("%+v\n", b.findAccount(id))
}

func (b *Bank) IncrementDate() {
	b.date = b.date + 1
	for _, a := range b.accountList {
		a.ApplyDailyInterest()
		if (b.date-a.GetdateOpen())%30 == 0 {
			a.UpdateOutstandingBalance()
		}
	}
}

func (b *Bank) IncrementDateBy(d int) {
	for i := 0; i < d; i++ {
		b.IncrementDate()
	}
}

package bank

import (
	"fmt"

	"../account"
)

type Bank struct {
	accountList []*account.Account
}

func (b *Bank) CreateAccount(Apr float32, CreditLimit float32) {
	var ID = len(b.accountList)
	var c = account.NewAccount(ID, Apr, CreditLimit)
	b.accountList = append(b.accountList, c)
}

func (b *Bank) findAccount(ID int) *account.Account {
	return b.accountList[ID]
}

func (b *Bank) ChargeAccount(ID int, c float32) {
	b.findAccount(ID).Charge(c)
}

func (b *Bank) PaymentAccount(ID int, c float32) {
	b.findAccount(ID).Pay(c)
}

func (b *Bank) PrintAccountInfo(ID int) {
	fmt.Printf("%+v\n", b.accountList[ID])
}

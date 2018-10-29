package bank

import (
	"fmt"

	"../account"
)

type Bank struct {
	accountList []*account.Account
	date        int
}

func (b *Bank) CreateAccount(apr float32, creditLimit float32) error {
	if apr <= 1 && apr >= 0 {
		var id = len(b.accountList)
		var c = account.NewAccount(id, apr, creditLimit, b.date)
		b.accountList = append(b.accountList, c)
	} else {
		return fmt.Errorf("APR (%g) requires to be between [0,1] ", apr)
	}
	return nil
}

func (b *Bank) findAccount(id int) (*account.Account, error) {
	for _, a := range b.accountList {
		if a.Getid() == id {
			return b.accountList[id], nil
		}
	}
	return nil, fmt.Errorf("account not found")
}

func (b *Bank) Charge(id int, c float32) error {
	a, err := b.findAccount(id)
	if err != nil {
		return err
	} else {
		if a.GetcreditLimit() > a.Getbalance()+c {
			a.Charge(c, b.date)
			return nil
		} else {
			return fmt.Errorf("over credit limit")
		}
	}
}

func (b *Bank) GetListOfCharges(id int) error {
	a, err := b.findAccount(id)
	if err != nil {
		return err
	} else {
		fmt.Printf("List of Charges: \n")
		for _, t := range a.GetcardSwipes() {
			fmt.Printf("Amount: %.2f		Date: %d\n", t.GetAmount(), t.GetDate())
		}
		return nil
	}
}

func (b *Bank) MakePayment(id int, c float32) error {
	a, err := b.findAccount(id)
	if err != nil {
		return err
	} else {
		a.Payment(c, b.date)
		return nil
	}
}

func (b *Bank) GetListOfPayments(id int) error {
	a, err := b.findAccount(id)
	if err != nil {
		return err
	} else {
		fmt.Printf("List of Payments: \n")
		for _, t := range a.Getpayments() {
			fmt.Printf("Amount: %.2f		Date: %d\n", t.GetAmount(), t.GetDate())
		}
		return nil
	}
}

func (b *Bank) GetOutstandingBalance(id int) error {
	a, err := b.findAccount(id)
	if err != nil {
		return err
	} else {
		fmt.Printf("OutstandingBalance: \n %.2f", a.GetoutstandingBalance())
		return nil
	}
}

func (b *Bank) Getinterest(id int) (float32, error) {
	a, err := b.findAccount(id)
	if err != nil {
		return 0, err
	}
	return a.Getinterest(), nil
}

func (b *Bank) PrintAccountInfo(id int) error {
	a, err := b.findAccount(id)
	if err != nil {
		return err
	} else {
		fmt.Printf("%+v\n", a)
		return nil
	}
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

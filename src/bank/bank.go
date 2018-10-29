package bank

import (
	"fmt"

	"../account"
)

// Bank represets list of creditcards and mangages date
type Bank struct {
	accountList []*account.Account
	date        int
}

// CreateAccount creates new account
func (b *Bank) CreateAccount(apr float32, creditLimit float32) (*account.Account, error) {
	if apr <= 1 && apr >= 0 {
		var id = len(b.accountList)
		var a = account.NewAccount(id, apr, creditLimit, b.date)
		b.accountList = append(b.accountList, a)
		return a, nil
	} else {
		return nil, fmt.Errorf("APR (%g) requires to be between [0,1]", apr)
	}
}

func (b *Bank) findAccount(id int) (*account.Account, error) {
	for _, a := range b.accountList {
		if a.Getid() == id {
			return b.accountList[id], nil
		}
	}
	return nil, fmt.Errorf("account not found")
}

// Charge apply a Charge with value c to account with id
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

// GetListOfCharges prints and returns the list of Charges([]*transaction.Transaction)
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

// MakePayment apply a Payment with value c to an account with id
func (b *Bank) MakePayment(id int, c float32) error {
	a, err := b.findAccount(id)
	if err != nil {
		return err
	} else {
		a.Payment(c, b.date)
		return nil
	}
}

// GetListOfPayments prints and returns the list of Payments([]*transaction.Transaction)
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

// GetOutstandingBalance print and returns the GetoutstandingBalance
func (b *Bank) GetOutstandingBalance(id int) (float32, error) {
	a, err := b.findAccount(id)
	if err != nil {
		return 0, err
	} else {
		fmt.Printf("OutstandingBalance: \n%.2f\n", a.GetoutstandingBalance())
		return a.GetoutstandingBalance(), nil
	}
}

// Getinterest print and returns the interest
func (b *Bank) Getinterest(id int) (float32, error) {
	a, err := b.findAccount(id)
	if err != nil {
		return 0, err
	}
	fmt.Printf("Interest: \n%.2f\n", a.Getinterest())
	return a.Getinterest(), nil
}

// IncrementDate increments the Date by one
func (b *Bank) IncrementDate() {
	b.date = b.date + 1
	for _, a := range b.accountList {
		a.ApplyDailyInterest()
		if (b.date-a.GetdateOpen())%30 == 0 {
			a.UpdateOutstandingBalance()
		}
	}
}

// IncrementDateBy increments the Date by d
func (b *Bank) IncrementDateBy(d int) {
	for i := 0; i < d; i++ {
		b.IncrementDate()
	}
}

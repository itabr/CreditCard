package bank

import (
	"testing"

	"../account"
)

func TestCreateAccount(t *testing.T) {
	var b = Bank{}
	_, test1 := b.CreateAccount(-0.1, 100)

	Expected := "APR (-0.1) requires to be between [0,1]"

	if test1.Error() != Expected {
		t.Errorf("Error got = %v, and Expected = %v.", test1.Error(), Expected)
	}

	_, test2 := b.CreateAccount(0, 100)
	if test2 != nil {
		t.Errorf("TestCreateAccount failed. Error got = %v, and Expected = %v.", test1.Error(), nil)
	}
}

func TestFindAccount(t *testing.T) {

	var a = account.NewAccount(0, 0.35, 1000, 0)

	var b = Bank{
		accountList: []*account.Account{a},
		date:        0}

	Expected := "account not found"
	_, test1 := b.findAccount(1)
	if test1.Error() != Expected {
		t.Errorf("Error got = %v, and Expected = %v.", test1.Error(), Expected)
	}

	_, test2 := b.findAccount(0)
	if test2 != nil {
		t.Errorf("TestCreateAccount failed. Error got = %v, and Expected = %v.", test1.Error(), nil)
	}
}

func TestCharge(t *testing.T) {
	var a = account.NewAccount(0, 0.35, 1000, 0)

	var b = Bank{
		accountList: []*account.Account{a},
		date:        0}

	Expected := "over credit limit"
	test1 := b.Charge(0, 5000.0)
	if test1.Error() != Expected {
		t.Errorf("TestCreateAccount failed. Error got = %v, and Expected = %v.", test1.Error(), nil)
	}

	test2 := b.Charge(0, 100.0)
	if test2 != nil {
		t.Errorf("TestCreateAccount failed. Error got = %v, and Expected = %v.", test2.Error(), nil)
	}
	if a.GetoutstandingBalance() != 100.0 || a.GetcardSwipes()[0].GetAmount() != 100.0 || a.GetcardSwipes()[0].GetDate() != 0 {
		t.Error("TestCharge failed")
	}
}
func TestMakePayment(t *testing.T) {
	var a = account.NewAccount(0, 0.35, 1000, 0)

	var b = Bank{
		accountList: []*account.Account{a},
		date:        0}

	b.MakePayment(0, 100.0)
	if a.GetoutstandingBalance() != -100.0 || a.Getpayments()[0].GetAmount() != 100.0 || a.Getpayments()[0].GetDate() != 0 {
		t.Error("TestCharge failed")
	}
}

func TestIncrementDate(t *testing.T) {
	var a = account.NewAccount(0, 0.35, 1000, 0)

	var b = Bank{
		accountList: []*account.Account{a},
		date:        0}

	b.Charge(0, 500.0)

	for i := 0; i < 15; i++ {
		b.IncrementDate()
	}

	b.MakePayment(0, 200.0)
	for i := 0; i < 10; i++ {
		b.IncrementDate()
	}

	b.Charge(0, 100.0)

	for i := 0; i < 5; i++ {
		b.IncrementDate()
	}

	OutstandingBalance, _ := b.GetOutstandingBalance(0)
	if OutstandingBalance != 411.9863 {
		t.Errorf("TestIncrementDate failed. Error got = %v, and Expected = %v.", OutstandingBalance, 411.99)
	}
}

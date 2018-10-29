package account

import (
	"reflect"
	"testing"
)

func TestNewAccount(t *testing.T) {
	var a = Account{
		id:          0,
		apr:         0.35,
		creditLimit: 1000.00,
		dateOpen:    0}

	if reflect.DeepEqual(NewAccount(0, 0.35, 1000.00, 0), a) {
		t.Error("TestNewAccount failed")
	}
}
func TestUpdateBalance(t *testing.T) {
	var a = Account{
		apr:         0.35,
		creditLimit: 1000.00}
	a.updateBalance(500)
	if a.outstandingBalance != 500 {
		t.Error("TestUpdateBalance failed")
	}
}

func TestCharge(t *testing.T) {
	var a = Account{
		apr:         0.35,
		creditLimit: 1000.00}
	a.Charge(500, 0)
	if a.outstandingBalance != 500 || a.cardSwipes[0].GetAmount() != 500 || a.cardSwipes[0].GetDate() != 0 {
		t.Error("TestCharge failed")
	}
}

func TestPayment(t *testing.T) {
	var a = Account{
		apr:         0.35,
		creditLimit: 1000.00}
	a.Payment(500, 0)
	if a.outstandingBalance != -500 || a.payments[0].GetAmount() != 500 || a.payments[0].GetDate() != 0 {
		t.Error("TestPayment failed")
	}
}

func TestApplyDailyInterest(t *testing.T) {
	var a = Account{
		apr:         0.35,
		creditLimit: 1000.00}
	a.Charge(500, 0)
	a.ApplyDailyInterest()

	if a.interest != 0.4794520548 {
		t.Error("TestApplyDailyInterest failed")
	}
}

func TestUpdateOutstandingBalance(t *testing.T) {
	var a = Account{
		apr:         0.35,
		creditLimit: 1000.00}

	a.Charge(500, 0)
	a.ApplyDailyInterest()
	a.UpdateOutstandingBalance()

	if a.interest != 0 || a.outstandingBalance != 500.4794520548 {
		t.Error("TestUpdateOutstandingBalance failed")
	}
}

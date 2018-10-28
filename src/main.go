package main

import (
	"./bank"
)

func main() {

	var b = bank.Bank{}
	b.CreateAccount(0.35, 1000.00)

	b.Charge(0, 500.0)
	b.PrintAccountInfo(0)

	b.IncrementDateBy(15)

	b.MakePayment(0, 200.0)
	b.PrintAccountInfo(0)

	b.IncrementDateBy(10)

	b.Charge(0, 100.0)
	b.PrintAccountInfo(0)

	b.IncrementDateBy(5)

	b.PrintAccountInfo(0)

}

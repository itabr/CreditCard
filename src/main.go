package main

import (
	"fmt"
	"os"

	"./bank"
)

func main() {

	var b = bank.Bank{}

	if err := b.CreateAccount(0.35, 1000.00); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if err := b.Charge(0, 500.0); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	b.PrintAccountInfo(0)

	b.IncrementDateBy(15)

	b.MakePayment(0, 200.0)
	b.PrintAccountInfo(0)

	b.IncrementDateBy(10)

	b.Charge(0, 100.0)
	b.PrintAccountInfo(0)

	b.IncrementDateBy(5)

	b.PrintAccountInfo(0)

	b.GetListOfCharges(0)

	b.GetOutstandingBalance(0)

}

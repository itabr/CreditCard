package main

import (
	"fmt"
	"os"

	"./bank"
)

func main() {

	var b = bank.Bank{}

	a, err := b.CreateAccount(0.35, 1000.00)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if err := b.Charge(a.Getid(), 500.0); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	b.IncrementDateBy(15)

	if err := b.MakePayment(a.Getid(), 200.0); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	b.IncrementDateBy(10)

	if err := b.Charge(a.Getid(), 100.0); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	b.IncrementDateBy(5)

	b.GetOutstandingBalance(a.Getid())

}

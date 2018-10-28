package main

import (
	"fmt"

	"./bank"
)

func main() {

	var b = bank.Bank{}
	b.CreateAccount(0.35, 1000.00)
	b.PrintAccountInfo(0)

	b.ChargeAccount(0, 500.0)
	b.ChargeAccount(0, 500.0)
	b.ChargeAccount(0, 900.0)
	b.PaymentAccount(0, 900.0)
	b.PrintAccountInfo(0)

	b.CreateAccount(0.35, 10000000.00)
	b.PrintAccountInfo(1)

	fmt.Println("Hello, Gopher!")
}

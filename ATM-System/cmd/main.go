package main

import (
	"fmt"

	"atm-system/internal/atm"
	"atm-system/internal/bankservice"
	"atm-system/internal/cashdispenser"
)

func main() {
	// cash
	denoms := []cashdispenser.Denomination{
		{Value: 100, Count: 10},
		{Value: 50, Count: 10},
		{Value: 20, Count: 10},
	}
	disp := cashdispenser.NewCoinChangeDispenser(denoms)

	// bank
	bank := bankservice.NewBankService()
	cardNum := bank.CreateAccount("1234", "sharath")

	// atm
	atm := atm.NewATM(bank, disp)

	fmt.Println("---- BALANCE CHECK ----")
	bal, _ := atm.CheckBalance(cardNum, "1234")
	fmt.Println("Balance:", bal)

	fmt.Println("\n---- WITHDRAW 380 ----")
	notes, err := atm.Withdraw(cardNum, "1234", 380)
	if err != nil {
		fmt.Println("Withdraw error:", err)
	} else {
		fmt.Println("Notes dispensed:", notes)
	}

	fmt.Println("\n---- BALANCE AFTER ----")
	bal, _ = atm.CheckBalance(cardNum, "1234")
	fmt.Println("Balance:", bal)

	fmt.Println("\n---- WRONG PIN ----")
	_, err = atm.Withdraw(cardNum, "9999", 100)
	fmt.Println("Expected error:", err)
}

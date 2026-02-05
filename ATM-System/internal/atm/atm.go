package atm

import (
	"atm-system/internal/bankservice"
	"atm-system/internal/cashdispenser"
	"errors"
	"fmt"
	"sync/atomic"
)

type ATM struct {
	BankService *bankservice.Bankservice
	CashDispenser *cashdispenser.CoinChangeDispenser
	TransactionsCounter int64
}

func NewATM(bank *bankservice.Bankservice, disp *cashdispenser.CoinChangeDispenser) *ATM {
	return &ATM{BankService:bank,CashDispenser: disp}
}


func (a *ATM) AuthenticateUser(cardNum string, pin string) error {
	return a.BankService.Authenticate(cardNum,pin)
	
}

func (a *ATM) CheckBalance(cardNum, pin string) (float64, error) {
	if err:=a.AuthenticateUser(cardNum, pin); err!=nil {
		return 0, errors.New("User Authentication Error")
	}
		balance, err := a.BankService.GetBalance(cardNum, pin)
	if err != nil {
		fmt.Println("[ATM] Balance fetch failed:", err)
		return 0, err
	}

	fmt.Println("[ATM] Current balance:", balance)
	return balance, nil
}


func (a *ATM) Deposit(cardNum, pin string, amount float64) error {
	if err:=a.AuthenticateUser(cardNum, pin); err!=nil {
		return errors.New("User Authentication Error")
	}

	if err := a.BankService.DepositMoney(cardNum, pin, amount); err != nil {
		return err
	}

	atomic.AddInt64(&a.TransactionsCounter, 1)
	return nil
}


func (a *ATM) Withdraw(cardNum, pin string, amount int) (map[int]int, error) {
	if err:=a.AuthenticateUser(cardNum, pin); err!=nil {
		return nil, errors.New("User Authentication Error")
	}
	balance, err := a.BankService.GetBalance(cardNum, pin)
	if err != nil {
		return nil, err
	}

	if float64(amount) > balance {
		return nil, errors.New("insufficient balance")
	}

	notes, err := a.CashDispenser.Dispense(amount)
	if err != nil {
		return nil, err
	}

	if err := a.BankService.WithDrawMoney(cardNum, pin, float64(amount)); err != nil {
		return nil, err
	}

	atomic.AddInt64(&a.TransactionsCounter, 1)
	return notes, nil
}

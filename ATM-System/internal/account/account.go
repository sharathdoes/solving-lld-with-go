package account

import (
	"atm-system/internal/utils"
	"errors"
	"sync"
)

type Card struct {
	CardNum string
	Pin     string
}



type Account struct {
	Username   string
	AccountNum string
	Balance    float64
	Cards      []Card
	Mu         sync.RWMutex
}

func NewAccount(username string) *Account {
	random := utils.GenerateRandom16Digit().String()
	return &Account{Username: username, AccountNum: random, Balance: 1999}
}

func NewCard(pin string) *Card {
	random:=utils.GenerateRandom16Digit().String()
	return &Card{CardNum: random, Pin:pin}
}



func (a *Account) Withdraw(amount float64) error {
	a.Mu.RLock()
	if a.Balance<amount {
		return errors.New("Insufficient Balance")
	}
	a.Mu.RUnlock()
	a.Mu.Lock()
	a.Balance-=amount
	 a.Mu.Unlock()
	return nil
}

func (a *Account) Deposit(amount float64) error {
	a.Mu.RLock()
	if a.Balance+amount>100000 {
		return errors.New("Can't deposit huge money")
	}
	a.Mu.RUnlock()
	a.Mu.Lock()
	a.Balance+=amount
	a.Mu.Unlock()
	return nil
}

func (c *Card) VerifyCard(pin string) error {
	if c.Pin != pin {
		return errors.New("incorrect PIN")
	}
	return nil
}
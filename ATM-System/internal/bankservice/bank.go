package bankservice

import (
	"atm-system/internal/account"
	"errors"
	"sync"
)

type Bankservice struct {
	cardAccounts map[*account.Card]*account.Account
	Accounts map[string]*account.Account
	Cards map[string]*account.Card
	mu sync.RWMutex
}

func NewBankService() *Bankservice{
	return &Bankservice{
		cardAccounts: make(map[*account.Card]*account.Account),
		Accounts:     make(map[string]*account.Account),
		Cards:     make(map[string]*account.Card),
	}
}

func (b *Bankservice) CreateAccount( pin string, username string) string {
	acc:=account.NewAccount(username)
	card:=account.NewCard(pin)
	b.cardAccounts[card]=acc
	b.Accounts[acc.AccountNum]=acc
		b.Cards[card.CardNum] = card
	return card.CardNum
}

func(b *Bankservice) AddCard(pin string, accountNumber string ){
	acc:=b.Accounts[accountNumber]
	card:=account.NewCard(pin)
	b.cardAccounts[card]=acc
		b.Cards[card.CardNum] = card

}

func(b *Bankservice)Authenticate(cardNum string, pin string) error {
	card, ok := b.Cards[cardNum]
	if !ok {
		return errors.New("Card is fake")
	}
	if err:=card.VerifyCard(pin); err!=nil {
		return errors.New("Card pin entered wrong")
	}
	return nil

}
func (b *Bankservice) GetBalance(cardNum, pin string) (float64, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	card, ok := b.Cards[cardNum]
	if !ok {
		return 0, errors.New("card not found")
	}

	acc, ok := b.cardAccounts[card]
	if !ok {
		return 0, errors.New("account not found for card")
	}

	acc.Mu.RLock()
	defer acc.Mu.RUnlock()

	return acc.Balance, nil
}

func (b *Bankservice) DepositMoney(cardNum string, pin string, amount float64) error {
	b.mu.RLock()
	defer b.mu.RUnlock()
	card:=b.Cards[cardNum]

	acc, ok := b.cardAccounts[card]
	if !ok {
		return errors.New("account not found for card")
	}

	acc.Mu.Lock()
	defer acc.Mu.Unlock()
	acc.Deposit(amount)

	return nil
}

func (b *Bankservice) WithDrawMoney(cardNum string, pin string, amount float64) error {
	b.mu.RLock()
	defer b.mu.RUnlock()
	card:=b.Cards[cardNum]

	if err:=card.VerifyCard(pin); err!=nil {
		return errors.New("Incorrect pin")
	}

	acc, ok := b.cardAccounts[card]
	if !ok {
		return errors.New("account not found for card")
	}

	acc.Mu.Lock()
	defer acc.Mu.Unlock()
	acc.Withdraw(amount)

	return nil
}
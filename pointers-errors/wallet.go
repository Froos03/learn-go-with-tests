package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(b Bitcoin) {
	w.balance += b
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(b Bitcoin) error {
	if w.balance < b {
		return ErrInsufficientFunds
	}
	w.balance -= b
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

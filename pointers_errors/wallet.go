package pointerserrors

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrorInsufficientFunds = errors.New("insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrorInsufficientFunds
	}
	w.balance -= amount

	return nil
}

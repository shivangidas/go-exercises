package main

import (
	"errors"
	"fmt"
)

// methods and pointers
type Bitcoin int
type Wallet struct {
	amount Bitcoin
}

func (w *Wallet) Deposit(money Bitcoin) {
	w.amount += money
}

func (w Wallet) Balance() Bitcoin {
	return w.amount
}

var ErrInsufficientFunds = errors.New("insufficient balance")

func (w *Wallet) Withdraw(money Bitcoin) error {
	if money > w.amount {
		return ErrInsufficientFunds
	}
	w.amount -= money
	return nil
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

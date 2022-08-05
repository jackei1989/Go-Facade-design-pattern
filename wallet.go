package main

import "fmt"

type Wallet struct {
	Balance int
}

func newWallet() *Wallet {
	return &Wallet{Balance: 0}
}

func (w *Wallet) credit(amount int) {
	w.Balance += amount
	fmt.Println("Wallet balance added successfully!")
}

func (w *Wallet) debit(amount int) error {
	if w.Balance > amount {
		return fmt.Errorf("Insufficient balance!")
	}
	w.Balance -= amount
	fmt.Println("Wallet balance debited successfully!")
	return nil
}

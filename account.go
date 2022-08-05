package main

import "fmt"

type Account struct {
	UserName string
}

func newAccount(userName string) *Account {
	return &Account{UserName: userName}
}

func (account *Account) checkAccount(userName string) error {
	if account.UserName != userName {
		return fmt.Errorf("UserName is not correct!")
	}
	return nil
}

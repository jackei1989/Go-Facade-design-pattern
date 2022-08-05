package main

import "fmt"

type Notification struct{}

func (n *Notification) sendWalletCreditNotification() {
	fmt.Println("send Wallet Credit Notification")
}

func (n *Notification) sendDebitNotification() {
	fmt.Println("send Debit Notification")
}
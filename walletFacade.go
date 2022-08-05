package main

import "fmt"

type WalletFacade struct {
	Account      *Account
	Wallet       *Wallet
	Notification *Notification
}

func newWalletFacade(userName string) *WalletFacade {
	fmt.Println("start create account")

	walletFacade := &WalletFacade{
		Account:      newAccount(userName),
		Wallet:       newWallet(),
		Notification: &Notification{},
	}

	fmt.Println("account created")
	return walletFacade
}

func (w *WalletFacade) addMoneyToWallet(userName string, amount int) error {
	fmt.Println("start add money to wallet")
	if err := w.Account.checkAccount(userName); err != nil {
		return err
	}
	w.Wallet.credit(amount)
	w.Notification.sendWalletCreditNotification()
	return nil
}

func (w *WalletFacade) debitMoneyFromWallet(userName string, amount int) error {
	fmt.Println("start debit money from wallet")
	if err := w.Account.checkAccount(userName); err != nil {
		return err
	}
	if err := w.Wallet.debit(amount); err != nil {
		return err
	}
	w.Notification.sendDebitNotification()
	return nil
}

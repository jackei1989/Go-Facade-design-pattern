package main


func main() {
	WalletFacade := newWalletFacade("jackei")
	WalletFacade.addMoneyToWallet("jackei", 100)
	WalletFacade.debitMoneyFromWallet("jackei", 50)
}
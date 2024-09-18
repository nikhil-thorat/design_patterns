package main

import (
	"fmt"
	"log"
)

const (
	CREDIT = "CREDIT"
	DEBIT  = "DEBIT"
)

type Account struct {
	name string
}

func newAccount(accountName string) *Account {
	return &Account{
		name: accountName,
	}
}

func (a *Account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("INCORRECT ACCOUNT NAME")
	}

	fmt.Println("ACCOUNT VERIFIED")
	return nil
}

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("AMOUNT CREDITED SUCCESSFULLY")
	return
}

func (w *Wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("NOT ENOUGH BALANCE")
	}

	w.balance = w.balance - amount
	fmt.Println("AMOUNT DEBITED SUCCESSFULLY")
	return nil
}

type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (s *SecurityCode) checkCode(code int) error {
	if s.code != code {
		return fmt.Errorf("INCORRECT SECURITY CODE")
	}

	fmt.Println("SECURITY CODE VERIFIED")
	return nil
}

type Notification struct{}

func (n *Notification) sendWalletCreditNotification() {
	fmt.Println("CREDIT NOTIFICATION SENT")
}

func (n *Notification) sendWalletDebitNotification() {
	fmt.Println("DEBIT NOTIFICATION SENT")
}

type Ledger struct{}

func (l *Ledger) makeEntry(accountID string, transactionType string, amount int) {
	fmt.Printf("LEDGER %s is %s for %d\n", accountID, transactionType, amount)
	return
}

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func newWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("CREATING NEW ACCOUNT")

	walletFacade := &WalletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: &Notification{},
		ledger:       &Ledger{},
	}

	fmt.Println("ACCOUNT CREATED")
	return walletFacade

}

func (w *WalletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("ADDING MONEY")

	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, CREDIT, amount)

	return nil
}

func (w *WalletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("DEDUCTING MONEY")

	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}

	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, DEBIT, amount)

	return nil
}

func main() {

	fmt.Println()

	walletFacade := newWalletFacade("GOBANK", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("GOBANK", 1234, 10)
	if err != nil {
		log.Fatalf("ERROR : %s\n", err)
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("GOBANK", 1234, 5)
	if err != nil {
		log.Fatalf("ERROR : %s\n", err)
	}

}

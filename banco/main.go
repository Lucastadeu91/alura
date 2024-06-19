package main

import (
	"fmt"

	"github.com/Lucastadeu91/alura/banco/acconts"
	"github.com/Lucastadeu91/alura/banco/clients"
)

func PayBill(account verifyAccount, billBalance float64) {
	account.WithdrawBalance(billBalance)
}

type verifyAccount interface {
	WithdrawBalance(balance float64) string
}

func main() {
	clientSilvia := clients.Client{Name: "Silvia", DigitalNumber: "123.123.123.54", Profession: "Dev"}
	currentAccountSilvia := acconts.CurrentAccounts{Client: clientSilvia, Agencynumber: 123, Acountnumber: 123}
	currentAccountSilvia.Deposit(100)
	savingsAccountSilvia := acconts.SavingsAccounts{Client: clientSilvia, Agencynumber: 123, Acountnumber: 123}
	savingsAccountSilvia.Deposit(500)

	PayBill(&savingsAccountSilvia, 60)

	fmt.Println(currentAccountSilvia.GetBalance())
	fmt.Println(savingsAccountSilvia.GetBalance())

}

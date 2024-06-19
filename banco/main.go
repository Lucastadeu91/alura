package main

import (
	"fmt"

	"github.com/Lucastadeu91/alura/banco/acconts"
	"github.com/Lucastadeu91/alura/banco/clients"
)

func main() {
	clientSilvia := clients.Client{Name: "Silvia", DigitalNumber: "123.123.123.54", Profession: "Dev"}
	accountSilvia := acconts.CurrentAccounts{Client: clientSilvia, Agencynumber: 123, Acountnumber: 123, Balance: 1235}
	fmt.Println(accountSilvia)

}

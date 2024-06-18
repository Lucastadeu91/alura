package main

import (
	"fmt"

	"github.com/Lucastadeu91/alura/banco/acconts"
	"github.com/Lucastadeu91/alura/banco/clients"
)

func main() {
	clientSilvia := clients.Client{"Silvia", "355.546.688.98", "dev"}
	accountSilvia := acconts.CurrentAccounts{clientSilvia, 123, 123, 1235.}
	fmt.Println(accountSilvia)

}

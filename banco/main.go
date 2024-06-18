package main

import (
	"fmt"

	"github.com/Lucastadeu91/alura/banco/clients"
	"github.com/Lucastadeu91/alura/banco/currentAccounts"
)

func main() {
	clientSilvia := clients.Client{"Silvia", "355.546.688.98", "dev"}
	accountSilvia := currentAccounts.CurrentAccounts{clientSilvia, 123, 123, 1235.}
	fmt.Println(accountSilvia)

}

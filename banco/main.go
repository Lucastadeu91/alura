package main

import (
	"fmt"

	c "github.com/Lucastadeu91/banco/accounts"
	"github.com/Lucastadeu91/banco/client"
)

func main() {
	contaDaSilvia := c.ContaCorrente{Client: client.Client{Name: "Silvia"}, Balance: 300}
	contaDaMichel := c.ContaCorrente{Client: client.Client{Name: "Michel"}, Balance: 100}
	status := contaDaSilvia.Transfer(200, &contaDaMichel)
	fmt.Println(status)
	fmt.Println(contaDaSilvia.Client.Name)
	fmt.Println(contaDaMichel)
}

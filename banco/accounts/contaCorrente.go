package contas

import "github.com/Lucastadeu91/banco/client"

type ContaCorrente struct {
	Client       client.Client
	Agencynumber int
	Acountnumber int
	Balance      float64
}
type ContaPoupanca struct {
	Client       client.Client
	Agencynumber int
	Acountnumber int
	Balance      float64
}

func (c *ContaCorrente) Deposit(deposit float64) (string, float64) {
	if deposit > 0 {
		c.Balance += deposit
		return "Deposito realizado com sucesso", c.Balance
	} else {
		return "Valor do deposito negativo", c.Balance
	}
}

func (c *ContaCorrente) WithdrawBalance(withdraw float64) string {
	if withdraw > 0 && withdraw <= c.Balance {
		c.Balance -= withdraw
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Transfer(transferValue float64, acountToBeTransfer *ContaCorrente) bool {
	if transferValue < c.Balance && transferValue > 0 {
		c.Balance -= transferValue
		acountToBeTransfer.Deposit(transferValue)
		return true
	} else {
		return false
	}
}

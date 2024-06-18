package currentAccounts

import "github.com/Lucastadeu91/alura/banco/clients"

type CurrentAccounts struct {
	Client       clients.Client
	Agencynumber int
	Acountnumber int
	Balance      float64
}

func (c *CurrentAccounts) Deposit(deposit float64) (string, float64) {
	if deposit > 0 {
		c.Balance += deposit
		return "Deposito realizado com sucesso", c.Balance
	} else {
		return "Valor do deposito negativo", c.Balance
	}
}

func (c *CurrentAccounts) WithdrawBalance(withdraw float64) string {
	if withdraw > 0 && withdraw <= c.Balance {
		c.Balance -= withdraw
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *CurrentAccounts) Transfer(transferValue float64, acountToBeTransfer *CurrentAccounts) bool {
	if transferValue < c.Balance && transferValue > 0 {
		c.Balance -= transferValue
		acountToBeTransfer.Deposit(transferValue)
		return true
	} else {
		return false
	}
}

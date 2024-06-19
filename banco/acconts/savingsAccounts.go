package acconts

import "github.com/Lucastadeu91/alura/banco/clients"

type CurrentAccounts struct {
	Client                                clients.Client
	Agencynumber, Acountnumber, Operation int
	balance                               float64
}

func (c *CurrentAccounts) Deposit(deposit float64) (string, float64) {
	if deposit > 0 {
		c.balance += deposit
		return "Deposito realizado com sucesso", c.balance
	} else {
		return "Valor do deposito negativo", c.balance
	}
}

func (c *CurrentAccounts) WithdrawBalance(withdraw float64) string {
	if withdraw > 0 && withdraw <= c.balance {
		c.balance -= withdraw
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *CurrentAccounts) Transfer(transferValue float64, acountToBeTransfer *CurrentAccounts) bool {
	if transferValue < c.balance && transferValue > 0 {
		c.balance -= transferValue
		acountToBeTransfer.Deposit(transferValue)
		return true
	} else {
		return false
	}
}

func (c *CurrentAccounts) GetBalance() float64 {
	return c.balance
}

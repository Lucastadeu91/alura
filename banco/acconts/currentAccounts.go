package acconts

import "github.com/Lucastadeu91/alura/banco/clients"

type SavingsAccounts struct {
	Client                     clients.Client
	Agencynumber, Acountnumber int
	balance                    float64
}

func (c *SavingsAccounts) Deposit(deposit float64) (string, float64) {
	if deposit > 0 {
		c.balance += deposit
		return "Deposito realizado com sucesso", c.balance
	} else {
		return "Valor do deposito negativo", c.balance
	}
}

func (c *SavingsAccounts) WithdrawBalance(withdraw float64) string {
	if withdraw > 0 && withdraw <= c.balance {
		c.balance -= withdraw
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *SavingsAccounts) Transfer(transferValue float64, acountToBeTransfer *SavingsAccounts) bool {
	if transferValue < c.balance && transferValue > 0 {
		c.balance -= transferValue
		acountToBeTransfer.Deposit(transferValue)
		return true
	} else {
		return false
	}
}

func (c *SavingsAccounts) GetBalance() float64 {
	return c.balance
}

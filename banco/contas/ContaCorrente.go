package contas

import (
	"GoOOProject/banco/clientes"
)

type ContaCorrente struct {
	Cliente       clientes.Cliente
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := valor > 0 && valor <= c.Saldo
	if podeSacar {
		c.Saldo -= valor
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo Insuficiente!!!"
	}
}

func (c *ContaCorrente) Depositar(valor float64) string {
	if valor > 0 {
		c.Saldo = c.Saldo + valor
		return "Depóstio efetuado com sucesso!!!"
	} else {
		return "Valor não permitido para depósito. Verifique e tente novamente."
	}
}

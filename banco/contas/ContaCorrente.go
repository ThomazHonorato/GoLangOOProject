package contas

import (
	"GoOOProject/banco/clientes"
)

type ContaCorrente struct {
	Cliente       clientes.Cliente
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

// Construtor para criar uma conta utilizando encapsulamento no saldo
func NovaContaCorrente(cliente clientes.Cliente, numeroAgencia int, numeroConta int, saldoInicial float64) ContaCorrente {
	return ContaCorrente{
		Cliente:       cliente,
		NumeroAgencia: numeroAgencia,
		NumeroConta:   numeroConta,
		saldo:         saldoInicial,
	}
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := valor > 0 && valor <= c.saldo
	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso!"
	} else {
		return "saldo Insuficiente!!!"
	}
}

func (c *ContaCorrente) Depositar(valor float64) string {
	if valor > 0 {
		c.saldo = c.saldo + valor
		return "Depóstio efetuado com sucesso!!!"
	} else {
		return "Valor não permitido para depósito. Verifique e tente novamente."
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}

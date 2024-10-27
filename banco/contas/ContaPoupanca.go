package contas

import (
	"GoOOProject/banco/clientes"
)

type ContaPoupanca struct {
	Cliente       clientes.Cliente
	NumeroConta   int
	NumeroAgencia int
	Operacao      int
	saldo         float64
}

func NovaContaPoupanca(cliente clientes.Cliente, numeroAgencia int, numeroConta int, saldoInicial float64) ContaCorrente {
	return ContaCorrente{
		Cliente:       cliente,
		NumeroAgencia: numeroAgencia,
		NumeroConta:   numeroConta,
		saldo:         saldoInicial,
	}
}

func (c *ContaPoupanca) Sacar(valor float64) string {
	podeSacar := valor > 0 && valor <= c.saldo
	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso!"
	} else {
		return "saldo Insuficiente!!!"
	}
}

func (c *ContaPoupanca) Depositar(valor float64) string {
	if valor > 0 {
		c.saldo = c.saldo + valor
		return "Depóstio efetuado com sucesso!!!"
	} else {
		return "Valor não permitido para depósito. Verifique e tente novamente."
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}

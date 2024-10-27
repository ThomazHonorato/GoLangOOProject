package main

import (
	"GoOOProject/banco/clientes"
	"GoOOProject/banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	pagamento := conta.Sacar(valorBoleto)
	if pagamento == "Saque realizado com sucesso!" {
		fmt.Printf("Pagamento realizado com sucesso!\n")
	} else {
		fmt.Printf("Erro ao efetuar pagamento!\n")
	}
}

type verificarConta interface {
	Sacar(valorBoleto float64) string
}

func main() {
	fmt.Printf("Sistema de contas\n\n")

	clienteThomaz := clientes.Cliente{
		Nome:      "Thomaz Honorato Soares",
		Cpf:       "12345678900",
		Profissao: "Desenvolvedor",
	}

	contaCorrenteThomaz := contas.NovaContaCorrente(
		clienteThomaz,
		108400,
		108400557,
		1500.00,
	)

	contaPoupancaThomaz := contas.NovaContaPoupanca(
		clienteThomaz,
		1082734,
		18273383,
		2000.00)

	fmt.Printf("Conta Corrente \n\n")
	fmt.Println("Titular da conta:", clienteThomaz.Nome)
	fmt.Println("CPF:", clienteThomaz.Cpf)
	fmt.Println("Profissão:", clienteThomaz.Profissao)
	fmt.Println("Numero da Agencia:", contaCorrenteThomaz.NumeroAgencia)
	fmt.Println("Numero da Conta:", contaCorrenteThomaz.NumeroConta)
	fmt.Printf("Saldo: R$ %.2f \n", contaCorrenteThomaz.ObterSaldo())
	fmt.Println()
	fmt.Printf("Conta Poupanca \n \n")
	fmt.Println("Titular da conta:", clienteThomaz.Nome)
	fmt.Println("CPF:", clienteThomaz.Cpf)
	fmt.Println("Profissão:", clienteThomaz.Profissao)
	fmt.Println("Numero da Agencia:", contaPoupancaThomaz.NumeroAgencia)
	fmt.Println("Numero da Conta:", contaPoupancaThomaz.NumeroConta)
	fmt.Printf("Saldo: R$ %2.f \n\n", contaPoupancaThomaz.ObterSaldo())

	fmt.Printf("Pagando Boleto:\n")
	PagarBoleto(&contaCorrenteThomaz, 2000)
	fmt.Printf("Saldo: %2.f \n", contaCorrenteThomaz.ObterSaldo())
}

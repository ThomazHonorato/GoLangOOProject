package main

import (
	"GoOOProject/banco/clientes"
	"GoOOProject/banco/contas"
	"fmt"
)

func main() {
	fmt.Println("Sistema de contas")

	clienteThomaz := clientes.Cliente{
		Nome:      "Thomaz Honorato Soares",
		Cpf:       "12345678900",
		Profissao: "Desenvolvedor",
	}

	contaThomaz := contas.NovaContaCorrente(
		clienteThomaz,
		108400,
		108400557,
		1500.00,
	)

	fmt.Println("Titular da conta:", clienteThomaz.Nome)
	fmt.Println("CPF:", clienteThomaz.Cpf)
	fmt.Println("Profissão:", clienteThomaz.Profissao)
	fmt.Println("Numero da Agencia:", contaThomaz.NumeroAgencia)
	fmt.Println("Numero da Conta:", contaThomaz.NumeroConta)
	fmt.Printf("Saldo: R$ %.2f \n", contaThomaz.ObterSaldo())
}

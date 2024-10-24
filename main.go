package main

import (
	"GoOOProject/banco/clientes"
	"GoOOProject/banco/contas"
	"fmt"
)

func main() {
	fmt.Println("Sistema de contas")

	contaThomaz := contas.ContaCorrente{
		Cliente: clientes.Cliente{
			Nome:      "Thomaz Honorato Soares",
			Cpf:       "12345678900",
			Profissao: "Desenvolvedor",
		},
		NumeroAgencia: 1023,
		NumeroConta:   1023400,
		Saldo:         1500.00,
	}
	fmt.Println("Titular da conta:", contaThomaz.Cliente.Nome)
	fmt.Println("CPF:", contaThomaz.Cliente.Cpf)
	fmt.Println("Profissão: ", contaThomaz.Cliente.Profissao)
	fmt.Println("Numero da Agencia:", contaThomaz.NumeroAgencia)
	fmt.Println("Numero da Conta:", contaThomaz.NumeroConta)
	fmt.Printf("Saldo:R$%.2f \n", contaThomaz.Saldo)
}

package main

import (
	"banco/cliente"
	"banco/contas"
	"fmt"
)

type verificarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valor float64) {
	conta.Sacar(valor)
}

func main() {

	var titular *cliente.Titular
	titular = new(cliente.Titular)
	titular.Nome = "Ana Silva"
	titular.Cpf = "123.456.789-00"
	titular.Profissao = "Desenvolvedora"

	var contaNil *contas.ContaCorrente
	contaNil = new(contas.ContaCorrente)
	contaNil.Titular = *titular
	contaNil.Agencia = 123
	contaNil.NumeroConta = 456789
	contaNil.Depositar(1000)
	fmt.Println(*contaNil)
	fmt.Println("2", contaNil)

	saldo := contaNil.ObterSaldo()
	fmt.Printf("Saldo atual: %.2f\n", saldo)

	PagarBoleto(contaNil, 20)

	saldo = contaNil.ObterSaldo()
	fmt.Printf("Saldo após pagar boleto: %.2f\n", saldo)
}

package contas

import (
	"banco/cliente"
	"fmt"
)

type ContaCorrente struct {
	Titular     cliente.Titular
	Agencia     int
	NumeroConta int
	saldo       float64
}

func DobrarNumeroSemPonteiro(x int) {
	x = x * 2
	fmt.Println("Dentro do escopo", x)
}

func DobrarNumeroComPonteiro(x *int) {
	*x = *x * 2
	fmt.Println("Dentro do escopo com ponteiro", *x)
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := valor > 0 && valor <= c.saldo
	if !podeSacar {
		return "saldo insuficiente"
	}
	c.saldo -= valor
	return "Saque realizado com sucesso"
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor <= 0 {
		return "Valor de depósito inválido", c.saldo
	}
	c.saldo += valor
	return "Depósito realizado com sucesso", c.saldo
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor <= 0 || valor > c.saldo {
		return false
	}
	c.saldo -= valor
	contaDestino.Depositar(valor)
	return true
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}

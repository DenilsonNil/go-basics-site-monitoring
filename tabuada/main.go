package main

import (
	"fmt"
)

const bar = "================================"

type Agente struct {
	Nome  string
	Idade int
}

func main() {

	fmt.Println(bar, "\nApp de Tabuada\n", bar)
	var numero int
	fmt.Print("Digite um número para ver a tabuada: ")
	fmt.Scan(&numero)

	fmt.Printf("Tabuada do %d:\n", numero)

	var agente *Agente
	agente = new(Agente)
	agente.Nome = "TabuadaAgent"
	agente.Idade = 1
	agente.process(numero)

	fmt.Println(agente.Nome, "processou a tabuada.")
	fmt.Println(agente)
	fmt.Println(*agente)
	fmt.Println(&agente)

}

func (*Agente) process(numero int) {
	for i := 1; i <= 10; i++ {
		resultado := numero * i
		fmt.Printf("%d x %d = %d\n", numero, i, resultado)
	}
}

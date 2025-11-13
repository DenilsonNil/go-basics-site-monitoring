package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Denilson"
	idade := 30
	altura := 1.75

	fmt.Println("Nome:", nome, "Tipo:", reflect.TypeOf(nome))
	fmt.Println("Idade:", idade, "Tipo:", reflect.TypeOf(idade))
	fmt.Println("Altura:", altura, "Tipo:", reflect.TypeOf(altura))
}

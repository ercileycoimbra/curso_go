package main

import "fmt"

func imprimirResultado(nota float64) {
	if (nota >= 7) || (nota < 2 && nota > 0) {
		fmt.Println("Aprovado com a nota", nota)
	} else {
		fmt.Println("Reprovado com a nota", nota)
	}
}

func main() {
	imprimirResultado(1.0)
	imprimirResultado(2.0)
}

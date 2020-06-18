package main

import "fmt"

// nao existe operador ternario em go
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "aprovado"
	}

	return "reprovado"
}

func main() {
	a := obterResultado(12.0)
	fmt.Println(a)
}

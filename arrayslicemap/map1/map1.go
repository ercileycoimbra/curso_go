package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123123123] = "Maria"
	aprovados[312321321] = "Ana"
	aprovados[456456456] = "Pedro"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[123123123])

	for id, _ := range aprovados {
		fmt.Println("ID:", id)
	}

}

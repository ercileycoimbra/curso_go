package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados:")
	for i, nome := range aprovados {
		fmt.Printf(" %d - %s\n", i+1, nome)
	}
}

func main() {
	aSlice := make([]string, 5)
	aSlice[0] = "Maria"
	aSlice[1] = "João"
	aSlice[2] = "Pedro"
	aSlice[3] = "Eco"
	aSlice[4] = "Laura"

	// a := []string{"Nome1", "Nome2"}

	imprimirAprovados(aSlice...)

}

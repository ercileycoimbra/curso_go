package main

import "fmt"

func main() {
	aNumeros := [...]int{12, 15, 9, 3, 56, 32} //compilador conta

	for i, numero := range aNumeros {
		fmt.Printf("Índice %d) = %d\n", i, numero)
	}

	for _, num := range aNumeros {
		fmt.Println(num)
	}
}

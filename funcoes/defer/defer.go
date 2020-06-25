package main

import "fmt"

func obterValor(n int) int {
	defer fmt.Println("Defer faz isso ser executado logo antes do return da função")

	if n >= 5000 {
		fmt.Printf("Valor máximo: %d\n", n)
		return 5000
	}

	fmt.Printf("Valor abaixo: %d\n", n)
	return n

}

func main() {
	fmt.Println(obterValor(2550))
	fmt.Println(obterValor(6000))
}

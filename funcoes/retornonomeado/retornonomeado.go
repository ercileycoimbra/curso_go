package main

import "fmt"

func troca(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return //você já definiu na declaração da função as variáveis que vão ser retornadas
}

func main() {
	a, b := troca(15, 27)
	fmt.Println(a, b)
}

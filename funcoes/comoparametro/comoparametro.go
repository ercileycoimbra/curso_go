package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	// funcao - parametro 1 da funcao exec
	// func - tipo do parametro 1
	// (int - parametro 1 da funcao passada como parametro
	// , int - parametro 2 da funcao passada como parametro
	// ) int - retorno da funcao passada como parametro
	// ,p1 - parametro p1 da funcao exec
	// , p2 - parametro p2 da funcao exec
	// int - tipo dos parametros p1 e p2
	// ) int - retorno da funcao exec

	return funcao(p1, p2)
}

func main() {
	res := exec(multiplicacao, 12, 4)
	fmt.Println(res)
}

package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) //está criando um slice referenciando 10 elementos de um array, mas o array será criado com 20 posições
	fmt.Println(s, len(s), cap(s))

	s = append(s, 3, 5, 6, 7, 9) //adiciona elementos no slice, mas ele continua apontando para o array criado pela função make, que tem capacidade de 20
	fmt.Println(s, len(s), cap(s))

	s = append(s, 3, 5, 6, 7, 9, 23)
	fmt.Println(s, len(s), cap(s))

}

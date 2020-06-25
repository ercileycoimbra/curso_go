package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		res, _ := fatorial(n - 1)
		return n * res, nil
	}
}

func main() {
	n, err := fatorial(-1)
	if err != nil {
		fmt.Printf("Erro, verifique: %s\n", err)
	} else {
		fmt.Printf("O fatorial deu %d\n", n)
	}
}

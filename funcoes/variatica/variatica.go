package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0

	for _, valor := range numeros {
		total += valor
	}

	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média do aluno foi de %.2f\n", media(1.4, 10.0, 8.0, 6.0))
}

package main

import "fmt"

func main() {
	var aNotas [3]float64

	aNotas[0] = 7.8
	aNotas[1] = 4.3
	aNotas[2] = 9.1

	fmt.Println(aNotas)

	var nMedia float64

	for AA := 0; AA < len(aNotas); AA++ {
		nMedia += aNotas[AA]
	}
	nMedia /= float64(len(aNotas))

	fmt.Printf("A média das notas foi %.2f\n", nMedia)
	fmt.Println((8.0 + 21.0 + (21 / 2)) * 9.0 * 67.05)

}

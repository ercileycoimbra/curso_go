package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(soma(2, 5), sub(10, 2))

}

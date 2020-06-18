package main

import "fmt"

func notaParaConceito(n float64) string {
	var c string

	if n >= 9 {
		c = "A"
	} else if n >= 8 {
		c = "B"
	} else if n >= 5 {
		c = "C"
	} else if n >= 3 {
		c = "D"
	} else {
		c = "E"
	}
	return c
}

func main() {
	cRes := notaParaConceito(8.2)

	fmt.Println("Nota para conceito =", cRes)
}

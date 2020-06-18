package main

import (
	"fmt"
)

func main() {
	fmt.Println("Nota para conceito:", notaParaConceito(8.7))
}

func notaParaConceito(n float64) string {
	var c string

	switch {
	case n >= 9:
		c = "A"
	case n >= 8:
		c = "B"
	case n >= 5:
		c = "C"
	case n >= 3:
		c = "D"
	default:
		c = "E"
	}
	return c
}

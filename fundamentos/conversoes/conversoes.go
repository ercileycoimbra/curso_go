package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	//dah erro fmt.Println(x / y)
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	cNota := fmt.Sprintf("%v", nota)

	fmt.Println(notaFinal, string(notaFinal), strconv.Itoa(123), cNota)

}

package main

import (
	"fmt"
)

func main() {
	var a = 11
	var b = 10
	fmt.Println("soma =>", a+b)
	fmt.Println("subtração =>", a-b)
	fmt.Println("divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Modulo (resto) =>", a%b)

	fmt.Println("E =>", a&b)
	fmt.Println("Ou =>", a|b)
	fmt.Println("Xor =>", a^b)

}

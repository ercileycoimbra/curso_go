package main

import "fmt"

func inc1(n int) {
	n++
}
func inc2(n *int) {
	*n++
}

func main() {
	a := 10
	inc1(a)
	fmt.Println(a)

	inc2(&a)
	fmt.Println(a)
}

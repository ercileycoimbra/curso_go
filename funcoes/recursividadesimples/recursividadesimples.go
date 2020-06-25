package main

import "fmt"

func newFat(n uint) uint {
	if n <= 1 {
		return 1
	}
	return n * newFat(n-1)
}

func main() {
	fmt.Println(newFat(5), newFat(3), newFat(0))
}

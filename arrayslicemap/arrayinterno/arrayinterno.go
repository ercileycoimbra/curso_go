package main

import (
	"fmt"
	"reflect"
)

func main() {
	aVet := [5]int{5, 10, 323, 65, 213}

	ad := make([]int, 1)
	ad = append(ad, 1)

	fmt.Println(ad, "ad")

	aSlice := aVet[0:5]

	aVet[1] += 32

	fmt.Println(reflect.TypeOf(aSlice), reflect.TypeOf(aVet))

	fmt.Println(aVet, aSlice)
}

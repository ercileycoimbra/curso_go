package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice

	fmt.Println(a1, s1)

	a2 := [5]int{1, 2, 3, 4, 5}

	//slice não é um array! Slice define um pedaço de um array
	s2 := a2[1:3]

	//alterando valor de um vetor
	a2[2] = 12
	fmt.Println(a2, s2)

	s3 := a2[:2] //novo slice, mas aponta para o mesmo array
	fmt.Println(a2, s3)

	//vc pode imaginar um slice como: tamanho e um ponteiro pra um elemento de um array
	s4 := s1[:1]
	fmt.Println(s2, s4)

}

package main

import "fmt"

func main() {
	i := 1
	i++

	var p *int = nil
	p = &i
	*p++
	i++

	//go não tem aritmética de ponteiros
	p++

	fmt.Println(p, *p, i, &i)
}

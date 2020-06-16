package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	nVal := 10000
	fmt.Println("A variável", nVal, "é do tipo:", reflect.TypeOf(nVal))

	//sem sinal, só positivos --- uint8 uint16 uint32 uint64
	var b byte = 255 //apelido para uint8
	fmt.Println("A variável", b, "é do tipo:", reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("A variável", i1, "é do tipo:", reflect.TypeOf(i1))

	fmt.Printf("A variável %v é do tipo %v\n", i1, reflect.TypeOf(i1))

	cTexto := "Teste com a crase"

	fmt.Printf("%s\nTipo é: %v", cTexto, reflect.TypeOf(cTexto))
}

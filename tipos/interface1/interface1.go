package main

import (
	"fmt"
	"reflect"
)

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

//interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	oPe := pessoa{"Ana", "Clara"}
	oPr := produto{"Bateria", 2.40}

	var oTeste imprimivel = pessoa{"Maria", "José"}

	imprimir(oPe)
	imprimir(oPr)
	imprimir(oTeste)

	fmt.Println(reflect.TypeOf(oPe), reflect.TypeOf(oPr), reflect.TypeOf(oTeste))

	//oPe = produto{} //dará erro
	oTeste = produto{} //não dá erro
}

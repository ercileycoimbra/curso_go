package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	oCar        carro
	turboLigado bool
	carro       //	campo anonimo (pseudo herança)
}

func main() {
	var c string

	f := ferrari{}
	f.oCar.nome = "F40"
	f.oCar.velocidadeAtual = 120
	f.turboLigado = true

	//possível acessar direto devido a "pseudo herança"
	f.nome = "F50"
	f.velocidadeAtual = 140

	if f.turboLigado {
		c = "ligado"
	} else {
		c = "desligado"
	}

	fmt.Printf("João comprou uma Ferrari %s, e o turbo está "+c+"\n", f.nome)
	fmt.Println(f)
}

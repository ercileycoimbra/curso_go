package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (obj pessoa) getNomePessoa() string {
	return obj.nome + " " + obj.sobrenome
}

func (o *pessoa) setNomePessoa(nomeCompleto string) {
	aRet := strings.Split(nomeCompleto, " ")
	o.nome = aRet[0]
	o.sobrenome = aRet[1]
}

func main() {
	var oPessoa pessoa
	oPessoa.nome = "Erciley"
	oPessoa.sobrenome = "Coimbra"
	fmt.Println(oPessoa.getNomePessoa())

	oPessoa.setNomePessoa("Laura Coimbra")
	fmt.Println(oPessoa.getNomePessoa())
}

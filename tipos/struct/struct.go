package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//Método: função com receiver (receptor)
func (obj produto) precoComDesconto() float64 {
	return obj.preco * (1 - (obj.desconto / 100))
}

func main() {
	var oProduto1 produto
	oProduto1.nome = "Ovo"
	oProduto1.preco = 1.30
	oProduto1.desconto = 5.0

	fmt.Printf("O produto %s, custa R$ %.2f, mas com o desconto de %.0f%% custará apenas %.2f.\n",
		oProduto1.nome, oProduto1.preco, oProduto1.desconto, oProduto1.precoComDesconto())

	var oT produto
	oT = produto{
		nome: "Teste",
	}
	fmt.Println(oT)

	oG := produto{"Nome", 10.25, 10.0}
	fmt.Println(oG)
}

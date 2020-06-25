package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
	oxi       int
}

type pedido struct {
	userID int
	itens  []item
}

func (obj pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range obj.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func (obj *pedido) addItem(oItem item) {
	var nLen int

	obj.itens = append(obj.itens, item{})
	nLen = len(obj.itens) - 1

	obj.itens[nLen].produtoID = oItem.produtoID
	obj.itens[nLen].qtde = oItem.qtde
	obj.itens[nLen].preco = oItem.preco
}

func main() {
	var oPedido pedido
	var oItem item

	//forma de inicializar a struct
	oPedido2 := pedido{
		userID: 2,
		itens: []item{
			item{produtoID: 1, qtde: 2, preco: 10.0},
			// item{2, 5, 12.0}, //dah pra fazer assim, só que inserindo uma nova propriedade irá fazer dar erro
		},
	}

	oPedido.userID = 1

	oItem.produtoID = 1
	oItem.qtde = 2
	oItem.preco = 12.10
	oPedido.addItem(oItem)

	oItem.produtoID = 2
	oItem.qtde = 1
	oItem.preco = 23.49
	oPedido.addItem(oItem)

	oItem.produtoID = 3
	oItem.qtde = 100
	oItem.preco = 3.13
	oPedido.addItem(oItem)

	fmt.Println(oPedido.valorTotal())

	fmt.Println(oPedido, "\n\n", oPedido2)
}

package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int
	Nome  string
	Preco float64
	Tags  []string
}

func main() {
	//struct para json
	p1 := produto{1, "Notebook", 1890.90, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	//json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta","preco":9.90,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1], p2)
}

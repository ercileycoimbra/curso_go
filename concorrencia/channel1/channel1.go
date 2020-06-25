package main

import "fmt"

func main() {
	ch := make(chan int, 1) //criando um canal de inteiros com buffer de 1

	ch <- 1 //enviando dados para um canal (escrita)
	<-ch    //recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)

}

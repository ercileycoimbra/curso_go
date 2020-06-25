package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	c <- 1
	fmt.Println("Após 1")
	c <- 2
	fmt.Println("Após 2")
	c <- 3
	fmt.Println("Após 3")
	c <- 4
	fmt.Println("Após 4")
	c <- 5
	fmt.Println("Após 5")
	c <- 6
	fmt.Println("Após 6")
}

func main() {
	c := make(chan int, 3)
	go rotina(c)

	time.Sleep(time.Second)
	fmt.Println(<-c)
}

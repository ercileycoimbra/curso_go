package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	canal := make(chan string)
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(time.Second)
			canal <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}
	}()
	return canal
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			select {
			case ret := <-entrada1:
				canal <- ret
			case ret := <-entrada2:
				canal <- ret
			}
		}
	}()
	return canal
}

func main() {
	c := juntar(falar("João"), falar("Maria"))
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}

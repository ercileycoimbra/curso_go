package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%2 == 0 {
			return false
		}
	}
	return true
}

func primo(qtdPrimos int, c chan int) {
	inicio := 2
	for i := 1; i <= qtdPrimos; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 50)
				break
			}
		}
	}
	close(c) //fechando canal para evitar deadlock
}

func main() {
	// fmt.Println(isPrimo(1999))
	// os.Exit(1)

	canal := make(chan int, 5)

	go primo(1000, canal)

	for primo := range canal {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("\nFim")
}

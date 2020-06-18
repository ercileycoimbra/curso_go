package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	for i <= 10 { //não tem while em go
		fmt.Println(i)
		i++
	}

	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)
	time.Sleep(time.Second * 3)

	var f float64 = 1

	for {
		fmt.Println("Oie forever", f)
		time.Sleep(time.Duration(1) * time.Second)
		break
	}

}

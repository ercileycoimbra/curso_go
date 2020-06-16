package main

import (
	"fmt"
	"math"
)

//import "fmt"

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo float64

	area := PI * math.Pow(raio, 2)

	fmt.Print("A área da bagaça é: ", area, round(area, 5))
}

func round(n float64, d float64) float64 {
	var a float64

	a = math.Trunc(n*math.Pow(10, d)) / math.Pow(10, d)
	return a
}

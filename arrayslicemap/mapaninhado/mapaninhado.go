package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   1200},
		"J": {
			"José João": 3123.12},
		"P": {
			"Pedro Junior": 5000},
	}

	for a, aMap := range funcsPorLetra {
		for nome, salario := range aMap {
			fmt.Println(a, nome, salario)
		}
	}
}

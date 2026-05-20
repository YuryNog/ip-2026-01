package main

import (
	"fmt"
	"math"
)

func main() {

	var vetor [15]float64
	var numero float64
	var i int

	for i = 0; i < 15; i++ {

		fmt.Print("Digite um número: ")
		fmt.Scan(&numero)

		if numero < 0 {
			vetor[i] = -1
		} else {
			vetor[i] = math.Sqrt(numero)
		}
	}

	fmt.Println("\nValores armazenados:")

	for i = 0; i < 15; i++ {
		fmt.Println(vetor[i])
	}
}
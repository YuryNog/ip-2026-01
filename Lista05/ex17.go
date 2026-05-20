package main

import "fmt"

func main() {

	var vetor [10]int
	var i int
	var j int
	var contador int

	for i = 0; i < 10; i++ {

		fmt.Printf("Digite o número %d: ", i+1)
		fmt.Scan(&vetor[i])
	}
	
	fmt.Println("\nNúmeros primos e suas posições:")

	for i = 0; i < 10; i++ {

		contador = 0

		for j = 1; j <= vetor[i]; j++ {

			if vetor[i]%j == 0 {
				contador++
			}
		}

		if contador == 2 {
			fmt.Println("Número:", vetor[i], "Posição:", i)
		}
	}
}
package main

import "fmt"

func main() {

	var vetor [10]int
	var i int
	var j int
	var aux int

	for i = 0; i < 10; i++ {

		fmt.Printf("Digite o número %d: ", i+1)
		fmt.Scan(&vetor[i])
	}

	for i = 0; i < 10; i++ {

		for j = i + 1; j < 10; j++ {

			if vetor[i] > vetor[j] {

				aux = vetor[i]
				vetor[i] = vetor[j]
				vetor[j] = aux
			}
		}
	}

	fmt.Println("\nVetor em ordem crescente:")

	for i = 0; i < 10; i++ {
		fmt.Print(vetor[i], " ")
	}
}
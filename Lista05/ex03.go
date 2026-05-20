package main

import "fmt"

func main() {

	var vetor [10]int
	var i int

	var somaPar int
	var quantImpar int

	for i = 0; i < 10; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&vetor[i])
	}

	fmt.Println("\nNúmeros pares:")

	for i = 0; i < 10; i++ {

		if vetor[i]%2 == 0 {
			fmt.Print(vetor[i], " ")

			somaPar = somaPar + vetor[i]
		}
	}

	fmt.Println("\n\nSoma dos números pares:")
	fmt.Println(somaPar)

	fmt.Println("\nNúmeros ímpares:")

	for i = 0; i < 10; i++ {

		if vetor[i]%2 != 0 {
			fmt.Print(vetor[i], " ")

			quantImpar++
		}
	}

	fmt.Println("\n\nQuantidade de números ímpares:")
	fmt.Println(quantImpar)
}
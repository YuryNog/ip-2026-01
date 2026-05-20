package main

import "fmt"

func main() {

	var vetor1 [10]int
	var vetor2 [10]int
	var resultado [20]int

	var i int
	var j int

	fmt.Println("Digite os valores do vetor 1:")

	for i = 0; i < 10; i++ {
		fmt.Scan(&vetor1[i])
	}

	fmt.Println("Digite os valores do vetor 2:")

	for i = 0; i < 10; i++ {
		fmt.Scan(&vetor2[i])
	}

	j = 0

	for i = 0; i < 10; i++ {

		resultado[j] = vetor1[i]
		j++

		resultado[j] = vetor2[i]
		j++
	}

	fmt.Println("\nVetor resultante:")

	for i = 0; i < 20; i++ {
		fmt.Print(resultado[i], " ")
	}
}
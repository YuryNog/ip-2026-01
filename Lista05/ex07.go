package main

import "fmt"

func main() {

	var vetor [100]int
	var i int
	var numero int

	numero = 1

	for i = 0; i < 100; i++ {

		vetor[i] = numero
		numero = numero + 2
	}

	fmt.Println("Números ímpares:")

	for i = 0; i < 100; i++ {
		fmt.Print(vetor[i], " ")
	}
}
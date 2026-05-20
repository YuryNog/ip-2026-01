package main

import "fmt"

func main() {

	var vetor [100]int
	var i int
	var numero int

	numero = 100

	for i = 0; i < 100; i++ {

		vetor[i] = numero
		numero--
	}

	fmt.Println("Elementos do vetor:")

	for i = 0; i < 100; i++ {
		fmt.Print(vetor[i], " ")
	}
}
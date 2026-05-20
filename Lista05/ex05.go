package main

import "fmt"

func main() {

	var vetor [10]int
	var i int

	var menor int
	var posicao int

	for i = 0; i < 10; i++ {
		fmt.Print("Digite um número: ")
		fmt.Scan(&vetor[i])
	}

	menor = vetor[0]
	posicao = 0

	for i = 0; i < 10; i++ {

		if vetor[i] < menor {
			menor = vetor[i]
			posicao = i
		}
	}

	fmt.Println("\nO menor elemento do vetor é:", menor)
	fmt.Println("E sua posição dentro do vetor é:", posicao)
}
package main

import "fmt"

func main() {

	var vetor [10]int
	var i int
	var soma int

	for i = 0; i < 10; i++ {
		fmt.Print("Digite um número: ")
		fmt.Scan(&vetor[i])
	}

	for i = 0; i < 10; i++ {

		if vetor[i] > 50 {
			fmt.Println("Número:", vetor[i])
			fmt.Println("Posição:", i)

			soma = soma + 1
		}
	}

	if soma == 0 {
		fmt.Println("Não existe número maior que 50")
	}
}
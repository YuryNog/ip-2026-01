package main

import "fmt"

func main() {

	var vetor [100]float64
	var i int

	var soma float64

	// Ler os valores
	for i = 0; i < 100; i++ {

		fmt.Printf("Digite o valor %d: ", i)
		fmt.Scan(&vetor[i])
	}

	// Calcular somatório
	for i = 0; i < 50; i++ {

		soma = soma + (vetor[i]-vetor[99-i])*3
	}

	// Mostrar resultado
	fmt.Println("\nValor do somatório:")
	fmt.Println(soma)
}
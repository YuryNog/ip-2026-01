package main

import "fmt"

func main() {

	var vetor1 [10]int
	var vetor2 [5]int

	var resultadoPar [10]int
	var resultadoImpar [10]int

	var i int
	var soma int
	var posPar int
	var posImpar int


	fmt.Println("Digite os valores do primeiro vetor:")

	for i = 0; i < 10; i++ {
		fmt.Printf("Valor %d: ", i+1)
		fmt.Scan(&vetor1[i])
	}

	fmt.Println("\nDigite os valores do segundo vetor:")

	for i = 0; i < 5; i++ {
		fmt.Printf("Valor %d: ", i+1)
		fmt.Scan(&vetor2[i])
	}

	for i = 0; i < 5; i++ {
		soma = soma + vetor2[i]
	}

	for i = 0; i < 10; i++ {

		if vetor1[i]%2 == 0 {
			resultadoPar[posPar] = vetor1[i] + soma
			posPar++
		} else {
			resultadoImpar[posImpar] = vetor1[i] + soma
			posImpar++
		}
	}

	fmt.Println("\nPrimeiro vetor resultante (pares):")

	for i = 0; i < posPar; i++ {
		fmt.Print(resultadoPar[i], " ")
	}

	fmt.Println("\n\nSegundo vetor resultante (ímpares):")

	for i = 0; i < posImpar; i++ {
		fmt.Print(resultadoImpar[i], " ")
	}
}
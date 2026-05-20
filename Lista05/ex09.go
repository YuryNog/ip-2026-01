package main

import "fmt"

func main() {

	var alturas [10]float64
	var i int

	var soma float64
	var media float64

	for i = 0; i < 10; i++ {

		fmt.Printf("Digite a altura do atleta %d: ", i+1)
		fmt.Scan(&alturas[i])

		soma = soma + alturas[i]
	}

	media = soma / 10


	fmt.Println("\nAlturas maiores que a média:")

	for i = 0; i < 10; i++ {

		if alturas[i] > media {
			fmt.Println(alturas[i])
		}
	}
}
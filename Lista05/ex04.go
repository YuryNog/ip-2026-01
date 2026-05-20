package main

import "fmt"

func main() {

	var A [10]int
	var i int
	var j int
	var contador int

	for i = 0; i < 10; i++ {
		fmt.Print("Digite um número: ")
		fmt.Scan(&A[i])
	}

	for i = 0; i < 10; i++ {

		contador = 0

		for j = 0; j < 10; j++ {

			if A[i] == A[j] {
				contador++
			}
		}

		if contador > 1 {
			fmt.Println("O número", A[i], "aparece", contador, "vezes")
		}
	}
}
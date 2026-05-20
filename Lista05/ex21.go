package main

import "fmt"

func main() {

	var vetor [10]float64
	var codigo int
	var i int

	fmt.Print("Digite o código: ")
	fmt.Scan(&codigo)

	if codigo == 0 {
		return
	}

	for i = 0; i < 10; i++ {

		fmt.Printf("Digite o valor %d: ", i+1)
		fmt.Scan(&vetor[i])
	}

	if codigo == 1 {

		fmt.Println("\nVetor na ordem direta:")

		for i = 0; i < 10; i++ {
			fmt.Print(vetor[i], " ")
		}
	}

	if codigo == 2 {

		fmt.Println("\nVetor na ordem inversa:")

		for i = 9; i >= 0; i-- {
			fmt.Print(vetor[i], " ")
		}
	}
}
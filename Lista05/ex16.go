package main

import "fmt"

func main() {

	var idades [50]int
	var i int
	var j int

	var contador int
	var maior int
	var moda int

	for i = 0; i < 50; i++ {

		fmt.Printf("Digite a idade %d: ", i+1)
		fmt.Scan(&idades[i])
	}

	for i = 0; i < 50; i++ {

		contador = 0

		for j = 0; j < 50; j++ {

			if idades[i] == idades[j] {
				contador++
			}
		}

		if contador > maior {
			maior = contador
			moda = idades[i]
		}
	}

	fmt.Println("\nModa das idades:", moda)
	fmt.Println("Quantidade de repetições:", maior)
}

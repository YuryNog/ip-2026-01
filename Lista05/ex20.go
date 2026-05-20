package main

import "fmt"

func main() {

	var dado [20]int
	var frequencia [7]int

	var i int

	for i = 0; i < 20; i++ {

		fmt.Printf("Digite o valor da jogada %d: ", i+1)
		fmt.Scan(&dado[i])

		frequencia[dado[i]]++
	}

	fmt.Println("\nNúmeros sorteados:")

	for i = 0; i < 20; i++ {
		fmt.Print(dado[i], " ")
	}

	fmt.Println("\n\nFrequência dos números:")

	for i = 1; i <= 6; i++ {

		fmt.Println("Número", i,
			"apareceu", frequencia[i], "vezes")
	}
}
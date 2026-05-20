package main

import "fmt"

func main() {

	var vetor1 [30]int
	var vetor2 [30]int

	var i int

	for i = 0; i < 30; i++ {

		fmt.Printf("Digite o valor %d: ", i+1)
		fmt.Scan(&vetor1[i])
	}

	for i = 0; i < 30; i++ {

		if i%2 == 0 {
			vetor2[i] = vetor1[i] * 2
		} else {
			vetor2[i] = vetor1[i] * 3
		}
	}

	fmt.Println("\nSegundo vetor:")

	for i = 0; i < 30; i++ {
		fmt.Print(vetor2[i], " ")
	}
}
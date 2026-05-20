package main

import "fmt"

func main() {

	var fibonacci [50]int
	var i int

	fibonacci[0] = 1
	fibonacci[1] = 1

	//Sequencia
	for i = 2; i < 50; i++ {

		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}

	fmt.Println("Sequência de Fibonacci:")

	for i = 0; i < 50; i++ {
		fmt.Print(fibonacci[i], " ")
	}
}
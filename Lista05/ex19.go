package main

import "fmt"

func main() {

	var num [10]int
	var divis [5]int

	var i int
	var j int

	fmt.Println("Digite os números do primeiro vetor:")

	for i = 0; i < 10; i++ {
		fmt.Scan(&num[i])
	}

	fmt.Println("Digite os números do segundo vetor:")

	for i = 0; i < 5; i++ {
		fmt.Scan(&divis[i])
	}

	for i = 0; i < 10; i++ {

		fmt.Println("\nNúmero", num[i], ":")

		for j = 0; j < 5; j++ {

			if num[i]%divis[j] == 0 {

				fmt.Println("Divisível por", divis[j],
					"na posição", j)
			}
		}
	}
}
//Exercicio 17

package main

import "fmt"

func main() {
	var x, y int

	fmt.Println("Digite dois numeros inteiros (x e y):")
	fmt.Scan(&x, &y)

	if x%2 != 0 {
		fmt.Println("O primeiro numero não é par")
	} else {
		fmt.Println("Sequencia de numeros pares:")

		for i := 0; i < y; i++ {
			fmt.Println(x + 2*i)
		}
	}
}

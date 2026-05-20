// Exercicio 15

package main

import "fmt"

func main() {
	var n int

	fmt.Println("Digite um valor inteiro:")
	fmt.Scan(&n)

	fmt.Println("Quadrados dos numeros pares de 1 ate", n, ":")

	for i := 1; i <= n; i++ {

		if i%2 == 0 {
			quadrado := i * i
			fmt.Println(i, "ao quadrado =", quadrado)
		}
	}
}

// Exercicio 10

package main

import "fmt"

func main() {
	var a, b, c, d float64
	var determinante float64

	fmt.Println("Digite o elemento a da matriz:")
	fmt.Scan(&a)

	fmt.Println("Digite o elemento b da matriz:")
	fmt.Scan(&b)

	fmt.Println("Digite o elemento c da matriz:")
	fmt.Scan(&c)

	fmt.Println("Digite o elemento d da matriz:")
	fmt.Scan(&d)

	// determinante
	determinante = a*d - b*c

	fmt.Printf("O valor do determinante e = %.2f\n", determinante)
}

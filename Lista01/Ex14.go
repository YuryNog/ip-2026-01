// Exercicio 14

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, h float64
	var Ab, V float64

	fmt.Println("Digite o comprimento da aresta do hexágono:")
	fmt.Scan(&a)

	fmt.Println("Digite a altura da pirâmide:")
	fmt.Scan(&h)

	Ab = (3 * a * a * math.Sqrt(3)) / 2

	V = (1.0 / 3.0) * Ab * h

	fmt.Printf("O volume da pirâmide e = %.2f\n", V)
}

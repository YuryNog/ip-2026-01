// Exercicio 8

package main

import "fmt"

func main() {

	var r float64
	var a float64
	var area float64
	var custo float64

	pi := 3.14159

	fmt.Println("Digite o raio da lata em centímetros: ")
	fmt.Scan(&r)

	fmt.Println("Digite a altura da lata em centímetros: ")
	fmt.Scan(&a)

	// Area do cilindro
	area = 2*(pi*r*r) + 2*pi*r*a

	// Custo
	custo = area * 0.01

	fmt.Printf("O valor do custo é = %.2f\n", custo)
}

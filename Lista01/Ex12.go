// Exercicio 12

package main

import "fmt"

func main() {
	var horas int
	var valor float64

	fmt.Println("Digite a quantidade de horas que a charrete foi usada:")
	fmt.Scan(&horas)

	blocos := horas / 3
	resto := horas % 3

	valor = float64(blocos*10) + float64(resto*5)

	fmt.Printf("O valor a pagar e = R$ %.2f\n", valor)
}

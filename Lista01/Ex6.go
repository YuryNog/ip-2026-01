// Exercicio 6

package main

import "fmt"

func main() {

	var F, C float64

	fmt.Println("Digite o valor em graus Fahrenheit: ")
	fmt.Scan(&F)

	C = 5 * (C - 32) / 9

	fmt.Printf("%.2f °F é igual a %.2f °C", F, C)
}

// Exercicio 7

package main

import "fmt"

func main() {

	var f float64
	var c float64
	var polegadas float64
	var mm float64

	// Temperatura
	fmt.Println("Digite a temperatura em Fahrenheit:")
	fmt.Scan(&f)

	c = (5*f - 160) / 9

	fmt.Println("Temperatura em Celsius:", c)

	// Chuva
	fmt.Println("Digite a quantidade de chuva em polegadas:")
	fmt.Scan(&polegadas)

	mm = polegadas * 25.4

	fmt.Println("Quantidade em milimetros:", mm)
}

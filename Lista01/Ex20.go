//Exercicio 20

package main

import "fmt"

func main() {
	var horas, minutos, segundos int
	var total int

	fmt.Println("Digite as horas:")
	fmt.Scan(&horas)

	fmt.Println("Digite os minutos:")
	fmt.Scan(&minutos)

	fmt.Println("Digite os segundos:")
	fmt.Scan(&segundos)

	// Converter
	total = horas*3600 + minutos*60 + segundos

	fmt.Println("Tempo total em segundos:", total)
}

// Exercicio 13

package main

import "fmt"

func main() {
	var nota float64
	var conceito string

	fmt.Println("Digite a nota do aluno (0 a 10):")
	fmt.Scan(&nota)

	if nota < 0 || nota > 10 {
		fmt.Println("Nota inválida")
		return
	}

	// Conversao conceito
	if nota >= 9.0 && nota <= 10.0 {
		conceito = "A"
	} else if nota >= 7.5 && nota < 9.0 {
		conceito = "B"
	} else if nota >= 6.0 && nota < 7.5 {
		conceito = "C"
	} else { // nota < 6.0
		conceito = "D"
	}

	fmt.Printf("O conceito correspondente e: %s\n", conceito)
}

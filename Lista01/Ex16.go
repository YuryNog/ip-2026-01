// Exercicio 16

package main

import "fmt"

func main() {
	var salario float64
	var novoSalario float64

	fmt.Println("Digite o salario atual:")
	fmt.Scan(&salario)

	if salario <= 300 {
		novoSalario = salario + (salario * 0.5)
	} else {
		novoSalario = salario + (salario * 0.3)
	}

	fmt.Println("Salario reajustado:", novoSalario)
}

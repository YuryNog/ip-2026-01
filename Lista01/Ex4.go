// Exercicio 4

package main

import "fmt"

func main() {

	var salario float64
	var kw float64

	fmt.Println("Digite o salario minimo:")
	fmt.Scan(&salario)

	fmt.Println("Digite a quantidade de kW:")
	fmt.Scan(&kw)

	// 100 kW = 70% do salario
	valor100 := salario * 0.7

	// valor de 1 kW
	valor1kw := valor100 / 100

	// total
	total := kw * valor1kw

	// desconto de 10%
	desconto := total * 0.1
	totalComDesconto := total - desconto

	fmt.Println("Valor de cada kW:", valor1kw)
	fmt.Println("Valor total:", total)
	fmt.Println("Valor com desconto:", totalComDesconto)
}

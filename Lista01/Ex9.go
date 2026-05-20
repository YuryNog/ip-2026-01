// Exercicio 9

package main

import "fmt"

func main() {
	var a, b, c float64
	var delta float64

	fmt.Println("Digite o valor de A: ")
	fmt.Scan(&a)

	fmt.Println("Digite o valor de B: ")
	fmt.Scan(&b)

	fmt.Println("Digite o valor de C: ")
	fmt.Scan(&c)

	//discriminante
	delta = b*b - 4*a*c

	fmt.Printf("O valor de delta é = %.2f\n", delta)
}

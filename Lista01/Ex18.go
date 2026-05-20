//Exercicio 18

package main

import "fmt"

func main() {
	var a1, r, n int
	var soma int
	var termo int

	fmt.Println("Digite o primeiro termo:")
	fmt.Scan(&a1)

	fmt.Println("Digite a razao:")
	fmt.Scan(&r)

	fmt.Println("Digite a quantidade de termos:")
	fmt.Scan(&n)

	soma = 0

	// Calcula termo por termo e soma
	for i := 0; i < n; i++ {
		termo = a1 + i*r
		fmt.Println("Termo", i+1, "=", termo)
		soma = soma + termo
	}

	fmt.Println("Soma dos termos =", soma)
}

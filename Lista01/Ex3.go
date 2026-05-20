// Exercicio 3

package main

import "fmt"

func main() {
	var n1, n2, n3 int

	fmt.Print("Digite tres números inteiros (n1 n2 n3): ")
	fmt.Scan(&n1, &n2, &n3)

	if n1 < 0 || n1 > 9 || n2 < 0 || n2 > 9 || n3 < 0 || n3 > 9 {
		fmt.Println("Invalido")
		return
	}

	//Concatenação
	numero := n1*100 + n2*10 + n3

	//Quadrado
	quadrado := numero * numero

	fmt.Println("Número:", numero)
	fmt.Println("Quadrado:", quadrado)
}

// Exercicio 11

package main

import "fmt"

func main() {

	var num int

	fmt.Println("Digite um número inteiro: \n")
	fmt.Scan(&num)

	if num%3 == 0 && num%5 == 0 {
		fmt.Println("O numero é divisel \n")
	} else {
		fmt.Println("o numero não e divisel \n")
	}
}

// Exercicio 2

package main

import "fmt"

func main() {
	var totalPublico float64
	var percPopular, percGeral, percArquibancada, percCadeiras float64

	fmt.Print("Digite o público total: ")
	fmt.Scan(&totalPublico)

	fmt.Print("Digite a porcentagem de Popular: ")
	fmt.Scan(&percPopular)

	fmt.Print("Digite a porcentagem de Geral: ")
	fmt.Scan(&percGeral)

	fmt.Print("Digite a porcentagem de Arquibancada: ")
	fmt.Scan(&percArquibancada)

	fmt.Print("Digite a porcentagem de Cadeiras: ")
	fmt.Scan(&percCadeiras)

	// Quantidade de pessoas por categoria
	qtdPopular := totalPublico * (percPopular / 100)
	qtdGeral := totalPublico * (percGeral / 100)
	qtdArquibancada := totalPublico * (percArquibancada / 100)
	qtdCadeiras := totalPublico * (percCadeiras / 100)

	// Valor
	valPopular := 1.0
	valGeral := 5.0
	valArquibancada := 10.0
	valCadeiras := 20.0

	// Lucro total
	rendaTotal := (qtdPopular * valPopular) +
		(qtdGeral * valGeral) +
		(qtdArquibancada * valArquibancada) +
		(qtdCadeiras * valCadeiras)

	fmt.Println("\nQuantidade por categoria:")
	fmt.Printf("Popular: %.0f\n", qtdPopular)
	fmt.Printf("Geral: %.0f\n", qtdGeral)
	fmt.Printf("Arquibancada: %.0f\n", qtdArquibancada)
	fmt.Printf("Cadeiras: %.0f\n", qtdCadeiras)

	fmt.Printf("\nRenda total: R$ %.2f\n", rendaTotal)
}

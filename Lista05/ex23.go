package main

import "fmt"

func main() {

	var janela [24]int
	var corredor [24]int

	var opcao int
	var poltrona int
	var i int
	var livresJanela int
	var livresCorredor int

	for {

		livresJanela = 0
		livresCorredor = 0

		for i = 0; i < 24; i++ {
			if janela[i] == 0 {
				livresJanela++
			}

			if corredor[i] == 0 {
				livresCorredor++
			}
		}

		if livresJanela == 0 && livresCorredor == 0 {
			fmt.Println("Ônibus completamente cheio.")
			break
		}

		fmt.Println("\n1 - Comprar poltrona na janela")
		fmt.Println("2 - Comprar poltrona no corredor")
		fmt.Println("3 - Finalizar")
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&opcao)

		if opcao == 1 {

			if livresJanela == 0 {
				fmt.Println("Não existem poltronas livres na janela.")
			} else {

				fmt.Println("Poltronas disponíveis na janela:")

				for i = 0; i < 24; i++ {
					if janela[i] == 0 {
						fmt.Print(i, " ")
					}
				}

				fmt.Print("\nEscolha a poltrona: ")
				fmt.Scan(&poltrona)

				if poltrona >= 0 && poltrona < 24 && janela[poltrona] == 0 {
					janela[poltrona] = 1
					fmt.Println("Passagem vendida.")
				} else {
					fmt.Println("Poltrona inválida ou ocupada.")
				}
			}
		}

		if opcao == 2 {

			if livresCorredor == 0 {
				fmt.Println("Não existem poltronas livres no corredor.")
			} else {

				fmt.Println("Poltronas disponíveis no corredor:")

				for i = 0; i < 24; i++ {
					if corredor[i] == 0 {
						fmt.Print(i, " ")
					}
				}

				fmt.Print("\nEscolha a poltrona: ")
				fmt.Scan(&poltrona)

				if poltrona >= 0 && poltrona < 24 && corredor[poltrona] == 0 {
					corredor[poltrona] = 1
					fmt.Println("Passagem vendida.")
				} else {
					fmt.Println("Poltrona inválida ou ocupada.")
				}
			}
		}

		if opcao == 3 {
			break
		}
	}
}
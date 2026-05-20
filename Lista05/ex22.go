package main

import "fmt"

func main() {

	var codigos [10]int
	var saldos [10]float64

	var i int
	var j int

	var menu int
	var codigo int
	var valor float64

	var encontrou bool
	var ativo float64

	for i = 0; i < 10; i++ {

		fmt.Printf("Digite o código da conta %d: ", i+1)
		fmt.Scan(&codigos[i])

		fmt.Printf("Digite o saldo da conta %d: ", i+1)
		fmt.Scan(&saldos[i])
	}

    	for {

		fmt.Println("\n1. Efetuar depósito")
		fmt.Println("2. Efetuar saque")
		fmt.Println("3. Consultar ativo bancário")
		fmt.Println("4. Finalizar programa")

		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&menu)

		if menu == 1 {

			fmt.Print("Digite o código da conta: ")
			fmt.Scan(&codigo)

			encontrou = false

			for j = 0; j < 10; j++ {

				if codigos[j] == codigo {

					fmt.Print("Digite o valor do depósito: ")
					fmt.Scan(&valor)

					saldos[j] = saldos[j] + valor

					fmt.Println("Depósito realizado!")

					encontrou = true
				}
			}

			if encontrou == false {
				fmt.Println("Conta não encontrada")
			}
		}

		if menu == 2 {

			fmt.Print("Digite o código da conta: ")
			fmt.Scan(&codigo)

			encontrou = false

			for j = 0; j < 10; j++ {

				if codigos[j] == codigo {

					fmt.Print("Digite o valor do saque: ")
					fmt.Scan(&valor)

					if saldos[j] >= valor {

						saldos[j] = saldos[j] - valor
						fmt.Println("Saque realizado!")

					} else {
						fmt.Println("Saldo insuficiente")
					}

					encontrou = true
				}
			}

			if encontrou == false {
				fmt.Println("Conta não encontrada")
			}
		}
		
		if menu == 3 {

			ativo = 0

			for j = 0; j < 10; j++ {
				ativo = ativo + saldos[j]
			}

			fmt.Println("Ativo bancário:", ativo)
		}
		
		if menu == 4 {
			break
		}
	}
}
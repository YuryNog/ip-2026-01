package main

import "fmt"

func main() {

	var codigo int
	var meses int
	var i int

	var cod1, cod2, cod3 int
	var mes1, mes2, mes3 int

	mes1 = 9999
	mes2 = 9999
	mes3 = 9999

	for i = 0; i < 100; i++ {

		fmt.Print("Digite o código do empregado: ")
		fmt.Scan(&codigo)

		fmt.Print("Digite os meses de trabalho: ")
		fmt.Scan(&meses)

		if codigo == 0 && meses == 0 {
			break
		}

		if meses < mes1 {

			mes3 = mes2
			cod3 = cod2

			mes2 = mes1
			cod2 = cod1

			mes1 = meses
			cod1 = codigo

		} else if meses < mes2 {

			mes3 = mes2
			cod3 = cod2

			mes2 = meses
			cod2 = codigo

		} else if meses < mes3 {

			mes3 = meses
			cod3 = codigo
		}
	}

	fmt.Println("\nEmpregados mais recentes:")

	fmt.Println("1º:", cod1, "-", mes1, "meses")
	fmt.Println("2º:", cod2, "-", mes2, "meses")
	fmt.Println("3º:", cod3, "-", mes3, "meses")
}

import "fmt"

func main() {

	var notas [15]int
	var frequencia [11]int

	var i int
	var relativa float64

	for i = 0; i < 15; i++ {

		fmt.Printf("Digite a nota do aluno %d: ", i+1)
		fmt.Scan(&notas[i])

		frequencia[notas[i]]++
	}

	// Mostrar tabela
	fmt.Println("\nNota | Frequência Absoluta | Frequência Relativa")

	for i = 0; i <= 10; i++ {

		relativa = float64(frequencia[i]) / 15

		fmt.Printf("%d     | %d                     | %.2f\n",
			i, frequencia[i], relativa)
	}
}
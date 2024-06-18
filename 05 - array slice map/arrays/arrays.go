package main

import "fmt"

func main() {
	// Arrays são homogêneos e estáticos (tam. fixo)
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	// notas[3] // Falha pois o array só posui 3 indexes de 0 a 2
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média é %.2f", media)
}

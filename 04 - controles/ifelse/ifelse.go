package main

import "fmt"

func imprimirResultado(nota float64) {
	// Não necessita de parêntesis, e as chaves são obrigatórias
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)
	imprimirResultado(4.9)
	imprimirResultado(8.1)
}

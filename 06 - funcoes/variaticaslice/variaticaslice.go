package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados!")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i, aprovado)
	}
}

func main() {
	aprovados := []string{"Marcos", "Aurélio", "Pedro"}
	// usando os pontos após o slice, ele é explodido
	imprimirAprovados(aprovados...)

	// é importante lembrar que o mesmo não pode ser feito com arrays,
	// portanto as duas hipóteses seguintes irão falhar

	// aprovadosArray := [3]string{"Marcos", "Aurélio", "Pedro"}
	// imprimirAprovados(aprovadosArray...)

	// aprovadosArray := [...]string{"Marcos", "Aurélio", "Pedro"}
	// imprimirAprovados(aprovadosArray...)
}

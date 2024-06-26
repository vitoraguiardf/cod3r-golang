package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// Error: assignment to entry in nil map
	// maps deve ser inicializados

	// inicializando map
	aprovados := make(map[int]string)

	// adcionando items ao map
	aprovados[12345678978] = "Maria"
	aprovados[12345687989] = "Pedro"
	aprovados[52845678952] = "Ana"

	fmt.Printf("aprovados: %v\n", aprovados)

	// percorrendo items do map
	for key, value := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", value, key)
	}

	// acessando item do map
	fmt.Printf("Aprovado: %v\n", aprovados[12345687989])

	// excluindo item do map
	delete(aprovados, 12345687989)
	fmt.Printf("aprovados: %v\n", aprovados)

	// sobrescrevendo item do map
	aprovados[52845678952] = "Pedro"
	fmt.Printf("aprovados: %v\n", aprovados)

}

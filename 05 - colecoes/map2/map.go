package main

import "fmt"

func main() {
	// inicializando map
	funcESalarios := map[string]float64{
		"José João":      11520.35,
		"Gabriela Silva": 15487.68,
		"Pedro Junior":   1200.80,
		// a vírgula é obrigatória ao final de cada item
		// error: unexpected newline in composite literal; possibly missing comma or }
	}
	fmt.Printf("funcESalarios: %v\n", funcESalarios)

	// adcionando item
	funcESalarios["Rafael Filho"] = 13500.72
	fmt.Printf("funcESalarios: %v\n", funcESalarios)

	// excluindo item inexistente, não gera erro
	delete(funcESalarios, "Inexistente")
	fmt.Printf("funcESalarios: %v\n", funcESalarios)

	// percorrenco itens do map
	for nome, salario := range funcESalarios {
		fmt.Printf("nome: %v >> salario: %v\n", nome, salario)
	}

	// percorrenco nomes do map
	for nome := range funcESalarios {
		fmt.Printf("nome: %v\n", nome)
	}

	// percorrenco salarios do map
	for _, salario := range funcESalarios {
		fmt.Printf("salario: %v\n", salario)
	}

}

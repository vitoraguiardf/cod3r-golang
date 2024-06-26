package main

import "fmt"

func main() {
	// iniciando map aninhado
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Germinni": 15456.78,
			"Guga Gole":         16547.85,
			"Genivaldo Gentil":  14586.48,
		},
		"J": {
			"José João":   5879.85,
			"Jonas Joule": 6547.526,
		},
		"P": {
			"Pedro Paulo":      12425.98,
			"Patrícia Paulada": 11256.881,
			"Poliana Pereira":  5869.45,
		},
	}
	fmt.Printf("funcsPorLetra: %v\n", funcsPorLetra)

	// acessando item(submap)
	fmt.Printf("funcsPorLetra[J]: %v\n", funcsPorLetra["J"])

	// removendo item(submap)
	delete(funcsPorLetra, "J")
	fmt.Printf("funcsPorLetra[J]: %v\n", funcsPorLetra["J"])

	// percorrendo items do main map
	for letra, funcs := range funcsPorLetra {
		fmt.Printf("funcs[%s]:\n", letra)
		// percorrendo items do sub map
		for name, salary := range funcs {
			fmt.Printf("name: %s, salary:%.3f\n", name, salary)
		}
	}

}

package main

import "fmt"

// Em Golang, não existe operador ternário
func obterResultado(nota float64) string {
	// A seguinte expressão comentada é inválida
	// return nota >= 6 ? "Aprovado" : "Reprovado"
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}
func main() {
	fmt.Println("Julia", obterResultado(6.2))
	fmt.Println("Paulo", obterResultado(6.0))
	fmt.Println("Clara", obterResultado(5.2))
	fmt.Println("Roberto", obterResultado(8.1))
}

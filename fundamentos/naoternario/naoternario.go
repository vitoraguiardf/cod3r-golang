package main

// Em Golang, não existe operador ternário
func obterResultado(nota float64) string {
	// A seguinte expressão comentada é inválida
	// return nota >= 6 ? "Aprovado" : "Reprovado"
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

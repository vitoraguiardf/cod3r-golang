package main

import (
	"fmt"
	"time"
)

// O parâmetro do tipo interface{} é um genérico que recebe qualquer tipo de valor
func tipo(i interface{}) string {
	// Usa se i.(type) para se obter o tipo da variável
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Function"
	default:
		return "Unknown"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Opa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}

package main

import "fmt"

func main() {
	fmt.Print("Impressão sem quebra de linha padrão, ")
	fmt.Print("deve ser passado manualmente com \\n.\n")

	fmt.Println("Imprime sempre com uma quebra de linha no final.")
	fmt.Println("Para cada chamada, uma nova linha.")

	const PI = 3.141592653589

	// fmt.Println("O valor de PI é: " + PI) 	// Não é possível concatenar dessa forma
	strPI := fmt.Sprint(PI)                  // retorna o valor como string
	fmt.Println("O valor de PI é: " + strPI) // É possível concatenar com mesmo tipo (string)

	fmt.Println("O valor de PI é", PI)      // Altenativa 1
	fmt.Printf("O valor de PI é: %.2f", PI) // Alternativa mais poderosa, permite formatação.

	a := 1
	b := 1.9999
	c := false
	d := "ôpa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d) // É possível formatar com patterns para vários tipos de variáveis
	fmt.Printf("\n%v %v %v %v", a, b, c, d)         // %v para saída de valores
	fmt.Printf("\n%T %T %T %T", a, b, c, d)         // %T para saída de tipos

}

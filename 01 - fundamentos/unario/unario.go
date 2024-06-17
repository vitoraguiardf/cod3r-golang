package main

import "fmt"

func main() {
	x, y := 1, 2

	// Apenas postfix
	x++ // Equivale a -> x += 1 ou x = x + 1
	fmt.Println(x)

	y-- // Equivale a -> x += 1 ou x = x + 1
	fmt.Println(y)

	// Não permitido
	// fmt.Println(x == y++)
}

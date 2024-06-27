package main

import "fmt"

func main() {
	x := 20
	fmt.Printf("x: %v (main func scope)\n", x)

	imprimeX := closure()
	imprimeX()
}

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Printf("x: %v (closure func scope)\n", x)
	}
	return funcao
}

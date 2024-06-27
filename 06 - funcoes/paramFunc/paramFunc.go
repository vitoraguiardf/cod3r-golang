package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func exec(operacao func(int, int) int, p1, p2 int) int {
	return operacao(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 3, 4)
	fmt.Printf("resultado: %v\n", resultado)
}

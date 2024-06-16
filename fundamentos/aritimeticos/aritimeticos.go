package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 2, 3

	fmt.Println("Aritmetics")
	fmt.Printf("Soma(%v, %v) => %v\n", a, b, (a + b))
	fmt.Printf("Subtração(%v, %v) => %v\n", a, b, (a - b))
	fmt.Printf("Divisão(%v, %v) => %v\n", a, b, (a / b))
	fmt.Printf("Multipicação(%v, %v) => %v\n", a, b, (a * b))
	fmt.Printf("Módulo(%v, %v) => %v\n", a, b, (a % b))

	fmt.Println("Bitwise")
	fmt.Printf("And(%v, %v) => %v\n", a, b, (a & b))
	fmt.Printf("Or(%v, %v) => %v\n", a, b, (a | b))
	fmt.Printf("Xor(%v, %v) => %v\n", a, b, (a ^ b))

	var c, d float64 = 2.0, 3.0
	// Pacote Math
	fmt.Printf("Max(%v, %v) => %v\n", a, b, math.Max(float64(a), float64(b)))
	fmt.Printf("Min(%v, %v) => %v\n", a, b, math.Min(c, d))
	fmt.Printf("Exponenciação(%v, %v) => %v\n", a, b, math.Pow(c, d))

}

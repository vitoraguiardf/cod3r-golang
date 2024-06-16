package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y)) // divisão de x diretamente por y não funciona por serem de tipos diferentes

	nota := 6.9
	notaFinal := int(nota)                   // conversão dec -> int somente ignora as casas decimais
	fmt.Println("A nota final e", notaFinal) // A nota final é 6

	// Cuidado!!!
	fmt.Println("O valor de string(97) é " + string(97)) // retorna o caractere correspondente

	// int -> string
	fmt.Println("A conversão de tipos int -> string é feita com strconv.Itoa(int).", strconv.Itoa(97))

	//string -> int
	num, _ := strconv.Atoi("123")
	fmt.Println("A conversão de tipos string -> int é feita com strconv.Atoi(int).", (num - 122))

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}

}

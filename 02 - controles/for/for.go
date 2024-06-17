package main

import (
	"fmt"
	"time"
)

func main() {
	// Em Golang não existe While, o For serve para todos os casos!

	// Loop for simulando um While
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Loop for tradicional
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	// Loop infinito
	for {
		fmt.Println("Forever...")
		time.Sleep(time.Second)
	}

	// Veremos o equivalente ao foreach na aula sobre arrays

}

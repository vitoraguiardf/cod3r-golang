package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5}
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	fmt.Println("Acessando somente indexes")
	for idx := range numeros {
		fmt.Println(idx)
	}

	fmt.Println("Acessando somente valores")
	for _, num := range numeros {
		fmt.Println(num)
	}
}

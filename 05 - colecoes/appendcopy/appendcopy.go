package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int // define uma coleção (array interno) com tamanho dinâmico
	fmt.Printf("array1: %v, slice1: %v\n", array1, slice1)

	// array1 = append(array1, 4, 5, 6)
	// Error: first argument to append must be a slice

	fmt.Println("Append [4, 5, 6] to slice1")
	slice1 = append(slice1, 4, 5, 6)
	fmt.Printf("array1: %v, slice1: %v\n", array1, slice1)

	slice2 := make([]int, 2) // define um slice de tamanho 2
	fmt.Printf("slice2: %v\n", slice2)

	fmt.Println("Copy slice1 to slice2")
	copy(slice2, slice1) // não aumenta o tamanho da coleção
	fmt.Printf("slice2: %v\n", slice2)
	// Copia somente os dois primeiros elementos devido a slice2 ter tamanho 2

}

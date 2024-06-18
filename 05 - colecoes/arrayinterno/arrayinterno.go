package main

import "fmt"

func main() {
	// Provando que slices, internamente, são referências e não cópias de um array
	// Criando um array
	a1 := [5]int{1, 2, 3, 4, 5}
	// Obtenco um slice desse array
	s1 := a1[:3]
	// Verificando
	fmt.Println(a1, s1)

	// Agora vamos alterar uma parte do slice
	s1[1] = 25
	// Verificando novamente
	fmt.Println(a1, s1)

	// Criando um slice de um slice
	s2 := s1[:1]
	// Verificando
	fmt.Println(a1, s1, s2)

	// Agora vamos alterar uma parte do novo slice
	s2[0] = 123
	// Verificando novamente
	fmt.Println(a1, s1, s2)

}

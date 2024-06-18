package main

import "fmt"

func main() {
	i := 1

	// Go não possui aritmetica de poteiros

	var p *int = nil // cria um ponteiro com endereço nulo

	p = &i // atribui um valor ao ponteiro

	fmt.Println(p, *p)
	fmt.Println(&i, i)

	i++
	fmt.Println(p, *p)
	fmt.Println(&i, i)

	*p--
	fmt.Println(p, *p)
	fmt.Println(&i, i)
}

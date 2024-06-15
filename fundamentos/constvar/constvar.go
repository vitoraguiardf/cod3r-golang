package main

import (
	"fmt"
	m "math" // possibilidade de renomear o nome do pacote importado
)

func main() {

	const PI float64 = 3.1415
	var raio = 3.2 // tipo float64 inferido pelo compilador

	// forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)
	// = atribui valor a variável
	// := para atribuição denâmica
	fmt.Println("A área da dirvunferência é", area)

	const (
		a = 1
		b = 2
		c = 3
	)

	var (
		d = 4
		e = 5
		f = 6
	)

	fmt.Println(a, b, c, d, e, f) // o compilador não permite que variaveis declaradas não seja utilizadas

	// atribuição respectiva de valores
	var g, h bool = true, false
	fmt.Println(g, h)

	// também funciona com tipos fiferentes
	i, j, k := 2, false, "êpa"
	fmt.Println(i, j, k)

}

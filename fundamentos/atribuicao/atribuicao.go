package main

import "fmt"

func main() {
	var b byte = 3 // Atribuição com tipo definido
	fmt.Println(b)

	i := 3 // Atribuição com inferência de tipo

	i += 3 // i = i + 3 -> Atribuição aditiva
	i -= 3 // i = i - 3 -> Atribuição subtrativa
	i /= 3 // i = i / 3 -> Atribuição divisiva
	i *= 3 // i = i * 3 -> Atribuição multiplicativa
	i *= 3 // i = i % 3 -> Atribuição modular

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}

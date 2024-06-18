package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Diferença ao declarar array e slice
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	// Slice não é um array! Slice define um pedaço de array.

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]

	fmt.Println(a2, s2)

	s3 := a2[:2] // um novo slice, mas aponta para um pedaço do array a2 declarado anteriormente
	fmt.Println(a2, s3)

	// Também é possível declararum slice a partir de outro slice
	s4 := s2[:1]
	fmt.Println(s2, s4)
}

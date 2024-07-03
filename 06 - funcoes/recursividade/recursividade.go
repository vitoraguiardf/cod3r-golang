package main

import "fmt"

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		anterior := fatorial(n - 1)
		return anterior * n
	}
}

func main() {
	res := fatorial(5)
	fmt.Printf("o resultado é %d\n", res)

	/* _, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}*/
}

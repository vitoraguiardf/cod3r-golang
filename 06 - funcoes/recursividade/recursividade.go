package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("o numero %d é inválido", n)
	case n == 0:
		return 1, nil
	default:
		anterior, _ := fatorial(n - 1)
		return anterior * n, nil
	}
}

func main() {
	res, _ := fatorial(5)
	fmt.Printf("o resultado é %d\n", res)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}

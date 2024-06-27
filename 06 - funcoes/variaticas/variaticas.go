package main

import "fmt"

func main() {
	media := media(7.7, 9.8, 8.1, 6.8)
	fmt.Printf("media: %.2f\n", media)
}

// é possível passar uma quantidade indeterminada de parametros
func media(nums ...float64) float64 {
	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total / float64(len(nums))
}

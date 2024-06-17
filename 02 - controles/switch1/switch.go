package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	// Em Golang, o switch executa o primeiro case válido e sai do bloco.
	// Esse comportamento pode ser alterado usando o comando falltrough
	switch nota {
	case 10:
		fallthrough // Continua para os proximos cases mesmo após executado o codigo deste case
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota Inválida!"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
	fmt.Println(notaParaConceito(-2.1))
	fmt.Println(notaParaConceito(11))
}

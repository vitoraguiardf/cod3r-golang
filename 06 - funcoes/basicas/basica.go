package main

import "fmt"

func main() {
	// invocando a função 1
	f1()

	// invocando a função 2
	f2("Valor 1", "Valor 2")

	// invocando a função 3
	f3response := f3()
	fmt.Printf("f3response: %s\n", f3response)

	// invocando a função 4
	f4response := f4("Valor 1", "Valor 2")
	fmt.Printf("f4response: %s\n", f4response)

	// recordando a atribuição de múltiplas variaveis
	f3res, f4res := f3(), f4("Ana", "Cecília")
	fmt.Printf("Respota de f3: %s, resposta de f4: %s\n", f3res, f4res)

	// invocando a função 5
	f5_r1, f5_r2 := f5()
	fmt.Printf("f5response: %s, %s\n", f5_r1, f5_r2)

}

// declaração básica de uma função
func f1() {
	fmt.Println("Executando a primeira função!")
}

// declaração de função com parâmetros
func f2(p1 string, p2 string) {
	fmt.Printf("Executando a segunda função! Parametros: p1: \"%s\", p2: \"%s\"\n", p1, p2)
}

// declaração de função com retorno
func f3() string {
	fmt.Println("Executando a terceira função!")
	return "\"Resposta da terceira função.\""
}

// declaração de função com declaração de parametos agrupada
func f4(p1, p2 string) string {
	return fmt.Sprintf("\"Resposta da quarta função. Parametros: p1: '%s', p2: '%s'\"", p1, p2)
}

// declaração de função retornando múltiplos valores
func f5() (string, string) {
	return "Verdadeiro", "Falso"
}

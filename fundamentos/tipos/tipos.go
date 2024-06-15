package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// signed (com bit de sinal)... int8 int16 int32 int64
	// unsigned (sem usar bit de sinal)... uint8 int16 int32 int64

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	fmt.Println("O valor máximo de int é", math.MaxInt)
	fmt.Println("O valor máximo de int8 é", math.MaxInt8)
	fmt.Println("O valor máximo de int16 é", math.MaxInt16)
	fmt.Println("O valor máximo de int32 é", math.MaxInt32)
	fmt.Println("O valor máximo de int64 é", math.MaxInt64)

	fmt.Println("O valor máximo de uint8 é", math.MaxUint8)
	fmt.Println("O valor máximo de uint16 é", math.MaxUint16)
	fmt.Println("O valor máximo de uint32 é", math.MaxUint32)

	var i2 rune = 'a'
	fmt.Println("O rune é do tipo", reflect.TypeOf(i2))
	fmt.Println("O valor é correspondente ao valor do caratere a na tabela unicode", i2)

	// float32 e float64
	var f float32 = 49.99
	fmt.Println("O tipo de f é ", reflect.TypeOf(f))
	fmt.Println("O valor de f é", f)
	fmt.Println("Já o tipo de 59.99 é", reflect.TypeOf(59.99))
	fmt.Println("Posso usar float32(59.99) para que o tipe seja", reflect.TypeOf(float32(59.99)))

	// boolean
	bo := true
	fmt.Println("O valoe de bo é", bo)
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println("Inversão com !bo resulta em", !bo)

	// string
	s1 := "Olá, meu nome é Vitor"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string s1 é", len(s1))

	// string com múltiplas linhas
	s2 := `Olá
	, meu nome é
	Vitor!
	`
	fmt.Println("O tamanho da string s2 é", len(s2))

	//char???
	char := 'b'
	fmt.Println("O valor de char é do tipo", reflect.TypeOf(char))
	fmt.Println("char retorna o id equivalente ao caractere na table unicode")
	fmt.Println("No caso, o valor de char = 'b' é", char)

}

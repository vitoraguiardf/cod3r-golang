package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Strings: \"Abacate\" == \"Abacate\"", "Abacate" == "Abacate")
	fmt.Println("Strings: \"Abacate\" == \"Banana\"", "Abacate" == "Banana")
	fmt.Println()

	fmt.Println("3 == 2", 3 == 2)
	fmt.Println("3 != 2", 3 != 2)

	fmt.Println("3 < 2", 3 < 2)
	fmt.Println("3 > 2", 3 > 2)

	fmt.Println("3 <= 2", 3 <= 2)
	fmt.Println("3 >= 2", 3 >= 2)
	fmt.Println()

	d1, d2 := time.Unix(0, 0), time.Unix(0, 0)
	fmt.Println("Datas: d1 == d2", d1 == d2)
	fmt.Println("Datas: d1.Equal(d2)", d1.Equal(d2))
	fmt.Println("Aqui a comparação foi feita com base em valores.")
	fmt.Println()

	type Pessoa struct {
		Nome string
	}

	p1, p2 := Pessoa{"Vitor"}, Pessoa{"Heitor"}
	fmt.Println("Pessoas: \"p1\" == \"p2\"", p1 == p2)
	fmt.Println("Aqui a comparação foi feita com base em valores.")
	fmt.Println()

}

# Desafio
## Estruturas de Controle
### if-else-if para switch

O desafio consiste em refatorar o seguinte código substituindo a estrutura if-else-if para switch.

```
func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
}
```

Resposta:

A estrutura ```switch {...``` é equivalente a ```switch true {...```, isso significa que será executado o bloco case em que o valor (ou resultado de expressão) seja igual a true.
```
func notaParaConceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}
}
```
package main

import "fmt"

// Escreva uma função que receba um mapa com valores inteiros e retorne a soma de todos os valores.

func somarValores(m map[string]int) int {
	soma := 0

	for _, valor := range m {
		soma += valor
	}

	return soma
}

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	total := somarValores(m)
	fmt.Println("A soma dos valores é:", total)
}

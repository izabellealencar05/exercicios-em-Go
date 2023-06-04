package main

import (
	"fmt"
	"strings"
)

// Escreva uma função que conte a ocorrência de cada palavra
// em uma string e retorne um mapa onde as chaves são as palavras
// encontradas e os valores são o número de ocorrências de cada palavra.

func countWords(s string) map[string]int {
	mapa := make(map[string]int)
	palavra := strings.Fields(s)

	for _, palavras := range palavra {
		mapa[palavras]++
	}
	return mapa
}
func main() {
	s := "olá, me chamo izabelle, e estou aprendendo Go"
	wordCount := countWords(s)

	for palavra, count := range wordCount {
		fmt.Printf("%s: %d\n", palavra, count)
	}
}

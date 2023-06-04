package main

import (
	"fmt"
	"sort"
)

func encontrarAnagramas(palavras []string) map[string][]string {
	anagramas := make(map[string][]string)

	for _, palavra := range palavras {

		sorted := sortString(palavra)

		anagramas[sorted] = append(anagramas[sorted], palavra)
	}

	return anagramas
}

func sortString(s string) string {

	bytes := []byte(s)

	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})

	return string(bytes)
}

func main() {
	palavras := []string{"clara", "giovana", "victor", "iza", "izabelle"}

	gruposAnagramas := encontrarAnagramas(palavras)

	for _, grupo := range gruposAnagramas {
		fmt.Println(grupo)
	}
}

package main

import "fmt"

// Escreva uma função que receba uma lista de mapas,
//onde cada mapa contém a contagem de palavras de um texto,
//e retorne um único mapa contendo a soma de todas as contagens.
func somarContagens(contagens []map[string]int) map[string]int {
	soma := make(map[string]int)

	for _, mapa := range contagens {
		for palavra, quantidade := range mapa {
			soma[palavra] += quantidade
		}
	}

	return soma
}

func main() {
	contagens := []map[string]int{
		{"bom": 2, "dia": 1},
		{"go": 3, "é": 1, "uma merda": 2},
	}

	somaContagens := somarContagens(contagens)

	for palavra, quantidade := range somaContagens {
		fmt.Printf("Palavra: %s, Quantidade: %d\n", palavra, quantidade)
	}
}
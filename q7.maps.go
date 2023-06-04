package main

import (
	"fmt"
	"strings"
)

//Escreva uma função que receba uma string contendo uma frase e retorne
//um mapa onde as chaves são as palavras encontradas na frase e os valores
//são mapas contendo a contagem de cada letra em cada palavra.

func contarLetrasPorPalavra(frase string) map[string]map[rune]int {
	palavras := obterPalavras(frase)
	contagem := make(map[string]map[rune]int)

	for _, palavra := range palavras {
		contagem[palavra] = contarLetras(palavra)
	}

	return contagem
}

func obterPalavras(frase string) []string {
	return strings.Fields(frase)
}

func contarLetras(palavra string) map[rune]int {
	contagem := make(map[rune]int)

	for _, letra := range palavra {
		contagem[letra]++
	}

	return contagem
}

func main() {
	frase := "A linguagem Go é incrível"

	contagemLetrasPorPalavra := contarLetrasPorPalavra(frase)

	for palavra, contagem := range contagemLetrasPorPalavra {
		fmt.Println("Palavra:", palavra)
		for letra, quantidade := range contagem {
			fmt.Printf("Letra: %c, Quantidade: %d\n", letra, quantidade)
		}
		fmt.Println()
	}
}
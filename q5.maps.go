package main

import "fmt"

//Escreva uma função que receba uma string e retorne um mapa onde as
//chaves são os caracteres presentes na string e os valores são a frequência de cada caractere

func contarCaracteres(str string) map[rune]int {
	frequencia := make(map[rune]int)

	for _, char := range str {
		frequencia[char]++
	}

	return frequencia
}

func main() {
	texto := "Go é um lixo, pior linguagem!"

	mapaFrequencia := contarCaracteres(texto)

	for char, freq := range mapaFrequencia {
		fmt.Printf("Caractere: %c, Frequência: %d\n", char, freq)
	}
}

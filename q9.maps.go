package main

import "fmt"

//Escreva uma função que gere a sequência de Fibonacci até
//um determinado número n e retorne um mapa onde as chaves
//são os índices da sequência e os valores são os números
//correspondentes.

func gerarSequenciaFibonacci(n int) map[int]int {
	sequencia := make(map[int]int)

	for i := 0; i <= n; i++ {
		if i <= 1 {
			sequencia[i] = i
		} else {
			sequencia[i] = sequencia[i-1] + sequencia[i-2]
		}
	}

	return sequencia
}

func main() {
	n := 10

	sequenciaFibonacci := gerarSequenciaFibonacci(n)

	for indice, numero := range sequenciaFibonacci {
		fmt.Printf("Índice: %d, Número: %d\n", indice, numero)
	}
}
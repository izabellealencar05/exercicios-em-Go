package main

import "fmt"

func contarPares(slice []int) map[[2]int]int {
	contagem := make(map[[2]int]int)

	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			pair := [2]int{slice[i], slice[j]}
			contagem[pair]++
		}
	}

	return contagem
}

func main() {
	numeros := []int{1, 2, 3, 4, 5, 1, 2, 3, 4}

	contagemPares := contarPares(numeros)

	for par, frequencia := range contagemPares {
		fmt.Printf("Par: %v, Frequência: %d\n", par, frequencia)
	}
}

package main

import (
	"fmt"
	"math"
)

// Crie um Slice de inteiros com tamanho 10 e imprima o valor mínimo e o valor máximo armazenados no Slice
func main() {
	slice := []int{10, 5, 3, 8, 2, 9, 7, 1, 6, 4}

	min := math.MaxInt64
	max := math.MinInt64

	for _, num := range slice {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	fmt.Println("Valor mínimo:", min)
	fmt.Println("Valor máximo:", max)

}

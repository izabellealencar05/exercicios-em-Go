package main

import "fmt"

// Crie um Array de floats com 10 elementos. Crie um novo Slice que contenha apenas os elementos do Array que são maiores que 5. Imprima o novo Slice.
func main() {
	array := [10]float64{2.3, 6.7, 1.2, 9.8, 4.5, 3.1, 7.6, 5.4, 8.2, 0.9}

	var slice []float64
	for _, num := range array {
		if num > 5 {
			slice = append(slice, num)
		}
	}

	fmt.Println("Novo slice contendo os elementos maiores que 5:", slice)
}


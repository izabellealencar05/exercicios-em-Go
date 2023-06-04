package main

import "fmt"

// Crie um Array de inteiros com 5 elementos. Em seguida, crie um novo Slice que contenha apenas os elementos do Array que são múltiplos de 3. Imprima o novo Slice.
func main() {
	array := [5]int{2, 6, 10, 15, 9}

	var slice []int
	for _, num := range array {
		if num%3 == 0 {
			slice = append(slice, num)
		}
	}

	fmt.Println("Novo slice contendo os elementos múltiplos de 3:", slice)
}

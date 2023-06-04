package main

import "fmt"

// Implemente uma função que receba um ponteiro para uma slice
// e seu tamanho e preencha-o com os n primeiros números primos.
func gerarPrimos(primos *[]int, n int) {
	if n <= 0 {
		return
	}

	*primos = append(*primos, 2)
	for i := 3; len(*primos) < n; i += 2 {
		primo := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				primo = false
				break
			}
		}
		if primo {
			*primos = append(*primos, i) // adiciona o número primo encontrado à slice
		}
	}
}

func main() {
	var primos []int
	gerarPrimos(&primos, 10)
	fmt.Println(primos)
}

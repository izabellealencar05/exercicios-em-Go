package main

import (
	"fmt"
	"math"
)

// Escreva uma função em Go que receba um ponteiro
// para uma variável float64 e atualize o valor da
// variável para a média aritmética entre o seu valor
// atual e o valor da constante Pi.
func atualizarMedia(valor *float64) {
	pi := math.Pi
	media := (*valor + pi) / 2
	*valor = media
}
func main() {
	var valor float64 = 3.14
	atualizarMedia(&valor)
	fmt.Println(valor)
}

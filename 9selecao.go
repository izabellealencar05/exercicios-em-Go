package main

import (
	"fmt"
	"sort"
)

// Algoritmo para ler três números reais e mostrá-los em ordem crescente:

func main() {
	var num1, num2, num3 float64

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)
	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&num3)

	numeros := []float64{num1, num2, num3}
	sort.Float64s(numeros)

	fmt.Println("Os números em ordem crescente são:", numeros)
}

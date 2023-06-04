package main

import "fmt"

// Algoritmo para imprimir a tabuada de multiplicação de 1 a 10 para um número fornecido pelo usuário:
func main() {
	var num int

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	for i := 1; i <= 10; i++ {
		fmt.Println(num, "x", i, "=", num*i)
	}

}

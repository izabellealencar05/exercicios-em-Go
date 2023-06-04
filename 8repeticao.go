package main

import "fmt"

// Algoritmo para imprimir todos os divisores de um número inteiro positivo fornecido pelo usuário:
func main() {
	var num int

	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&num)

	fmt.Println("Divisores de", num, ":")

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Println(i)
		}
	}
}

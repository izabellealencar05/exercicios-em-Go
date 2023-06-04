package main

import "fmt"

// Algoritmo para verificar se um número inteiro é múltiplo de 3 e de 5 ao mesmo tempo:
func main() {
	var num int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&num)

	if num%3 == 0 && num%5 == 0 {
		fmt.Println("O número é múltiplo de 3 e de 5 ao mesmo tempo.")
	} else {
		fmt.Println("O número não é múltiplo de 3 e de 5 ao mesmo tempo.")
	}
}

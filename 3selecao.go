package main

import "fmt"

// Faça um algoritmo que leia um número inteiro e mostre se ele é par ou ímpar.
func main() {
	var num int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("O número é par.")
	} else {
		fmt.Println("O número é ímpar.")
	}
}

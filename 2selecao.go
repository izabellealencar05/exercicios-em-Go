package main

import "fmt"

// Faça um algoritmo que leia três números inteiros e mostre o menor deles.
func main() {
	var num1, num2, num3 int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)
	fmt.Print("Digite o terceiro número: ")
	fmt.Scan(&num3)

	if num1 < num2 && num1 < num3 {
		fmt.Println("O menor número é:", num1)
	} else if num2 < num1 && num2 < num3 {
		fmt.Println("O menor número é:", num2)
	} else if num3 < num1 && num3 < num2 {
		fmt.Println("O menor número é:", num3)
	} else {
		fmt.Println("Os números são iguais.")
	}
}

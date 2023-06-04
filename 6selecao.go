package main

import "fmt"

// Algoritmo para multiplicar dois números inteiros se ambos forem positivos ou somá-los se pelo menos um deles for negativo:
func main() {
	var num1, num2 int

	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)

	if num1 >= 0 && num2 >= 0 {
		multiplicacao := num1 * num2
		fmt.Println("A multiplicação dos números é:", multiplicacao)
	} else {
		soma := num1 + num2
		fmt.Println("A soma dos números é:", soma)
	}
}

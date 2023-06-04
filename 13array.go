package main

import "fmt"

// Crie um Array de inteiros com 7 elementos. Solicite ao usuário que informe um número que será adicionado ao primeiro e ao último elemento do Array. Imprima o Array resultante.
func main() {

	array := [7]int{2, 5, 8, 11, 14, 17, 20}

	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	array[0] += numero
	array[len(array)-1] += numero

	fmt.Println("Array resultante:", array)
}


package main

import "fmt"

//Faça um programa que leia dois arrays de inteiros de tamanho n e gere um terceiro array que seja a soma dos dois primeiros.
func main() {

	var n int
	fmt.Print("Digite o tamanho dos arrays: ")
	fmt.Scanln(&n)

	array1 := make([]int, n)
	array2 := make([]int, n)
	array3 := make([]int, n)

	fmt.Println("Digite os elementos do primeiro array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scanln(&array1[i])
	}

	fmt.Println("Digite os elementos do segundo array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scanln(&array2[i])
	}

	for i := 0; i < n; i++ {
		array3[i] = array1[i] + array2[i]
	}

	fmt.Println("Terceiro array (soma):", array3)
}


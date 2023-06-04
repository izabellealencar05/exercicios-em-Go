package main

import "fmt"

// Algoritmo para ler vários números inteiros e mostrar a média aritmética entre eles. A leitura deve ser interrompida quando for lido o valor zero:
func main() {
	var num, count int
	sum := 0

	for {
		fmt.Print("Digite um número inteiro (digite 0 para sair): ")
		fmt.Scan(&num)

		if num == 0 {
			break
		}

		sum += num
		count++
	}

	if count > 0 {
		media := float64(sum) / float64(count)
		fmt.Println("Média:", media)
	} else {
		fmt.Println("Nenhum número foi digitado.")
	}
}

package main

import "fmt"

// Algoritmo para ler vários números inteiros e mostrar o maior deles. A leitura deve ser interrompida quando for lido o valor zero:
func main() {
	var num, maior int

	for {
		fmt.Print("Digite um número inteiro (digite 0 para sair): ")
		fmt.Scan(&num)

		if num == 0 {
			break
		}

		if num > maior {
			maior = num
		}
	}

	if maior != 0 {
		fmt.Println("O maior número digitado foi:", maior)
	} else {
		fmt.Println("Nenhum número foi digitado.")
	}
}

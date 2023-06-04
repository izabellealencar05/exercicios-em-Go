package main

import "fmt"

// Crie um Array bidimensional de inteiros com 2 linhas e 3 colunas. Em seguida, solicite ao usuário que informe um índice de linha e outro de coluna e imprima o valor armazenado nessa posição da matriz.
func main() {
	matriz := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	var linha, coluna int
	fmt.Print("Digite o índice da linha: ")
	fmt.Scan(&linha)
	fmt.Print("Digite o índice da coluna: ")
	fmt.Scan(&coluna)

	if linha >= 0 && linha < 2 && coluna >= 0 && coluna < 3 {

		valor := matriz[linha][coluna]
		fmt.Println("O valor armazenado na posição", linha, ",", coluna, "é:", valor)
	} else {
		fmt.Println("Índices inválidos!")
	}
}

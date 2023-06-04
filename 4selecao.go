package main

import "fmt"

// Algoritmo para verificar se uma pessoa está abaixo, dentro ou acima do peso ideal com base na altura e sexo:
func main() {
	var altura, pesoIdeal float64
	var sexo string

	fmt.Print("Digite a altura da pessoa em metros: ")
	fmt.Scan(&altura)
	fmt.Print("Digite o sexo da pessoa (M para masculino ou F para feminino): ")
	fmt.Scan(&sexo)

	if sexo == "M" {
		pesoIdeal = (72.7 * altura) - 58
	} else if sexo == "F" {
		pesoIdeal = (62.1 * altura) - 44.7
	} else {
		fmt.Println("Sexo inválido.")
		return
	}

	var peso float64
	fmt.Print("Digite o peso da pessoa: ")
	fmt.Scan(&peso)

	if peso < pesoIdeal {
		fmt.Println("A pessoa está abaixo do peso ideal.")
	} else if peso > pesoIdeal {
		fmt.Println("A pessoa está acima do peso ideal.")
	} else {
		fmt.Println("A pessoa está no peso ideal.")
	}
}

package main

import "fmt"

// Algoritmo para calcular o novo salário de um funcionário com base no salário atual e aplicar um aumento de 10% se o salário for menor que R$ 1000,00 ou de 5% se o salário for maior ou igual a R$ 1000,00:
func main() {
	var salario float64

	fmt.Print("Digite o salário do funcionário: ")
	fmt.Scan(&salario)

	var novoSalario float64
	if salario < 1000.00 {
		novoSalario = salario + (salario * 0.10)
	} else {
		novoSalario = salario + (salario * 0.05)
	}

	fmt.Println("O novo salário do funcionário é:", novoSalario)
}

package main

import "fmt"

//Escreva uma função que receba um mapa onde as chaves são nomes de pessoas
//e os valores são as quantias de dinheiro que cada pessoa gastou em uma conta
//compartilhada. A função deve calcular o valor que cada pessoa deve receber ou
//pagar para igualar as despesas.

func equalizarContas(gastos map[string]float64) map[string]float64 {
	totalGastos := calcularTotalGastos(gastos)
	numPessoas := float64(len(gastos))
	valorPorPessoa := totalGastos / numPessoas

	resultado := make(map[string]float64)

	for pessoa, gasto := range gastos {
		resultado[pessoa] = valorPorPessoa - gasto
	}

	return resultado
}

func calcularTotalGastos(gastos map[string]float64) float64 {
	total := 0.0

	for _, gasto := range gastos {
		total += gasto
	}

	return total
}

func main() {
	gastos := map[string]float64{
		"Alice":  50.0,
		"Bob":    30.0,
		"Carlos": 70.0,
		"David":  40.0,
	}

	resultado := equalizarContas(gastos)

	for pessoa, valor := range resultado {
		if valor > 0 {
			fmt.Printf("%s deve receber: R$ %.2f\n", pessoa, valor)
		} else if valor < 0 {
			fmt.Printf("%s deve pagar: R$ %.2f\n", pessoa, -valor)
		} else {
			fmt.Printf("%s não precisa receber nem pagar\n", pessoa)
		}
	}
}
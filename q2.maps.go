package main

//Escreva uma função que receba dois mapas e retorne um novo
//mapa contendo todos os elementos dos mapas de entrada.
//Em caso de chaves duplicadas, o valor do segundo mapa deve prevalecer.
import "fmt"

func mapas(mapa1, mapa2 map[string]int) map[string]int {
	ajuntar := make(map[string]int)

	for chave, valor := range mapa1 {
		ajuntar[chave] = valor
	}

	for chave, valor := range mapa2 {
		ajuntar[chave] = valor
	}

	return ajuntar
}

func main() {
	mapa1 := map[string]int{
		"chave1": 1,
		"chave2": 2,
	}

	mapa2 := map[string]int{
		"chave2": 3,
		"chave3": 4,
	}

	ajuntar := mapas(mapa1, mapa2)

	for chave, valor := range ajuntar {
		fmt.Printf("%s: %d\n", chave, valor)
	}
}

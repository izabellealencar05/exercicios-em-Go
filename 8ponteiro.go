package main

import "fmt"

// Crie uma função que receba um ponteiro para uma variável como
// argumento e modifique o valor da variável dentro da função.
// Certifique-se de que o ponteiro aponte para uma área de memória
// válida e que a memória seja liberada após o uso.

func modificarValor(p *int, novoValor int) {
	*p = novoValor
}
func main() {
	valor := 10
	ptr := &valor

	modificarValor(ptr, 20)

	fmt.Println(valor)

	ptr = nil
}

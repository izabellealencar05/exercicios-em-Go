package main

import "fmt"

// Escreva uma função em Go que receba um ponteiro
// para uma variável inteira e atualize o valor da
// variável para a soma dos valores dos seus dois
// últimos dígitos (por exemplo, se o valor inicial
// da variável for 1234, o novo valor será 3+4=7).
func ptr(n *int) {
	ultimosDigitos := *n % 100
	soma := (ultimosDigitos / 10) + (ultimosDigitos % 10)
	*n = soma
}
func main() {
	var numero int
	fmt.Println("escreva um numero com mais de 4 algarismos:")
	fmt.Scan(&numero)
	ptr(&numero)
	fmt.Println(numero)
}

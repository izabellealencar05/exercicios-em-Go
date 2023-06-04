package main

import "fmt"

// Escreva uma função em Go que receba
// um ponteiro para um struct Conta com
// campos saldo e titular, e adicione um
// valor ao saldo da conta.

type Conta struct {
	saldo   float64
	titular string
}

func ptr(conta *Conta, valor float64) {
	conta.saldo += valor
}
func main() {
	minhaConta := Conta{saldo: 100.0, titular: "izabelle"}
	ptr(&minhaConta, 60.0)
	fmt.Println("o nome do titular é", minhaConta.titular, "e seu saldo é:", minhaConta.saldo)
}

package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel2++
	fmt.Println(variavel1, variavel2)

	fmt.Println("-----------")

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int // * define que a variavel é um ponteiro

	variavel3 = 100
	ponteiro = &variavel3 // & cria a conexão

	fmt.Println(variavel3, ponteiro) // será impresso o valor em memória

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) // será impresso o valor da variável em memória
}

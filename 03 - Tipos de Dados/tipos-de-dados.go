package main

import (
	"errors"
	"fmt"
)

func main() {
	// NÚMEROS INTEIROS
	var numero int64 = -1000000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000 // uint só aceita números inteiros sem sinais
	fmt.Println(numero2)

	var numero3 rune = 123456 // rune = alias para int32
	fmt.Println(numero3)

	var numero4 byte = 123 // byte = alias para uint8
	fmt.Println(numero4)

	// NÚMEROS REAIS
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char) // Vai imprimir o valor ASCII, no caso A = 65

	texto := 5 // Vai entender que texto é um número inteiro
	fmt.Println(texto)

	var booleano bool
	fmt.Println(booleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}

package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua A", 123}

	usuario2 := usuario{"Maria", 32, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 24} // Para não precisar declarar todos os dados precisa especificar qual está sendo preenchido
	fmt.Println(usuario3)
}

package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco1 := enderecos.TipoDeEndereco("Avenida Paulista")
	tipoEndereco2 := enderecos.TipoDeEndereco("Rodovia dos Imigrantes")
	tipoEndereco3 := enderecos.TipoDeEndereco("Praça das Rosas")
	fmt.Println(tipoEndereco1)
	fmt.Println(tipoEndereco2)
	fmt.Println(tipoEndereco3)
}
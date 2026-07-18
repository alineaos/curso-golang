package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	// As Struct Tags não são obrigatórias, porém o JSON usará o nome igual o do atributo, mantendo as letras maiúsculas (se houver)
	// Usar "-" se não quiser que o campo apareça
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil{
		log.Fatal(erro)
	}

	fmt.Println()
	fmt.Println(cachorroEmJSON) // retorna um slice de byte/uint8
	fmt.Println(bytes.NewBuffer(cachorroEmJSON)) // retorna no formato JSON

	c2 := map[string]string{
		"nome": "Toby",
		"raca":"Poodle",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil{
		log.Fatal(erro)
	}
	fmt.Println()
	fmt.Println(cachorro2EmJSON) // retorna um slice de byte/uint8
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON)) // retorna no formato JSON
}

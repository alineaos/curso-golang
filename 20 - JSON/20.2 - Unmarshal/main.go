package main

import (
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
	cachorroEmJSON := `{"nome":"Rex","raca":"Dálmata","idade":3}`

	var c cachorro

	// Unmarshal precisa receber um slice de byte e o endereço em memória da variável que receberá o valor convertido
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)


	cachorro2EmJSON := `{"nome":"Toby","raca":"Poodle"}`

	c2 := make(map[string]string)


	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool{
return u.idade >= 18
}

func (u *usuario) fazerAniversario(){
	u.idade++
}

func main(){
	var usuario1 usuario = usuario{"Yuuji Iadori", 15}
	usuario1.salvar()

	usuario2 := usuario{"Satoru Gojou", 28}
	usuario2.salvar()

	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
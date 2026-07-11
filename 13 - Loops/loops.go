package main

import (
	"fmt"
	// "time"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"João", "Davi", "Lucas"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for i, letra := range "PALAVRA"{
		fmt.Println(i, string(letra))
	}

	usuario := map[string]string{
		"nome": "Leonardo",
		"sobrenome": "Silva",
	}

	for k, v := range usuario{
		fmt.Println(k,v)
	}
}

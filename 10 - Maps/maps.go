package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Yuuji",
		"sobrenome": "Itadori",
	}

	fmt.Println(usuario["nome"])
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Megumi",
			"ultimo":   "Fushiguro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Tokyo",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Capricórnio",
	}

	fmt.Println(usuario2)
}

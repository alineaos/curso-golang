package main

import "fmt"

func generica(inter interface{}) {
	fmt.Println(inter)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	// o Print é uma função com tipo genérico!
	fmt.Println(1, 2, "string", false, true, float64(12345))

	// mapa := map[interface{}]interface{}{
	// 	1:            "String",
	// 	float32(100): true,
	// 	"String":     "String",
	// 	true:         float64(12),
	// }

	// fmt.Println(mapa)
}
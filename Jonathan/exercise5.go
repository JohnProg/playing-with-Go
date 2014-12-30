package main

import "fmt"

func main(){
	elementos := map[string] map[string] string{
		"H": map[string] string{
			"Nombre": "Hidrogeno",
			"Estado": "Gas",
		},
		"O": map[string] string {
			"Nombre": "Oxigeno",
			"Estado": "Liquido",
		},
	}

	fmt.Println(elementos["H"]["Nombre"])
}
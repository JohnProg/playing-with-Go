package main

import "fmt"

func main(){
	x := make(map[string] int)
	fmt.Println("Tamaño", len(x))
	x["Edad"] = 24

	fmt.Println(x["Edad"])
	fmt.Println("Tamaño", len(x))
	delete(x, "Edad")
	fmt.Println("Tamaño", len(x))
}
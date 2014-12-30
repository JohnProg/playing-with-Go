package main

import "fmt"

func main(){
	var input string

	elements := make(map[string] string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println("Ingresa un simbolo")
	fmt.Scanf("%v", &input)

	if elemento, status:= elements[input]; status{
		fmt.Println("El elemento es: ", elemento)
	} else {
		fmt.Println("No existe el elemento")
	}



}
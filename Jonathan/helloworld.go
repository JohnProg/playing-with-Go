package main

import "fmt"

const constante float32 = 3.14

func main(){
	var nombre string = "Jonathan"
	var apellido string = "Carrasco"
	var input float32

	sobrenombre := "Jonas"
	fmt.Println("Hello World:", nombre == apellido, sobrenombre)
	fmt.Println("Mi constante es:", constante)

	fmt.Println("Ingrese un numero: ")
	fmt.Scanf("%f", &input)

	output := input * 2
	output +=1 

	fmt.Println("El resultado es: ", output)
}
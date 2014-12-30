package main

import "fmt"

func main(){
	var numeros[3] int

	for i:=0;i < len(numeros); i++{
		numeros[i] = i*2
	}

	/*
	for i:=0;i<len(numeros);i++{
		fmt.Println(numeros[i])
	}*/

	var total float64 = 0

	for _, value := range numeros {
		total += float64(value)
	}
	fmt.Println(total / float64(len(numeros)))
	 
}
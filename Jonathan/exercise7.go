package main

import "fmt"

func firts(){
	fmt.Println("Primero")
}

func second(){
	fmt.Println("Segundo")
}

func main(){
	defer second()
	firts()
	firts()	
	firts()
}
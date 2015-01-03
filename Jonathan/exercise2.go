package main

import "fmt"

func main(){
	x:= make([]int, 5, 10)
	fmt.Println(x)	
	for i:=0;i<len(x);i++{
		fmt.Println(x[i])
	}

}

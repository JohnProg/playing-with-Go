package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct{
	Stuff Data
}

type Data struct {
	Fruit Fruits
	Veggies Vegetables
}

type Fruits map[string] int
type Vegetables map[string] int 

func serveRest(w http.ResponseWriter, r *http.Request){
	response, error := getJsonResponse()
	if error != nil {
		panic(error)
	}
	fmt.Fprintf(w, string(response))
}

func main(){
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:1337", nil)
}

func getJsonResponse() ([]byte, error) {
	fruits :=make(map[string] int)
	fruits["Apples"] = 25
	fruits["Oranges"] =  11

	vegetables :=make(map[string] int)
	vegetables["Carrots"] = 21
	vegetables["Peppers"] = 8

	d := Data{fruits, vegetables}
	p :=Payload{d}

	return json.Marshal(p)

}
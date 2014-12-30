package main

import "fmt"
import "encoding/json"
import "net/http"
import "io/ioutil"

type Payload struct{
	Stuff Data
}

type Data struct {
	Fruit Fruits
	Veggies Vegetables
}

type Fruits map[string] int
type Vegetables map[string] int 

func main(){
	url := "http://localhost:1337/"
	res, error := http.Get(url)
	if error != nil{
		panic(error)
	}
	defer res.Body.Close()

	body, error := ioutil.ReadAll(res.Body)
	if error != nil{
		panic(error)
	}

	var p Payload

	error = json.Unmarshal(body, &p)
	if error !=nil{
		panic(error)
	}

	fmt.Println(p.Stuff.Fruit)
}
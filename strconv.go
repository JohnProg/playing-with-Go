package main

 import "fmt"
 import "strconv"


 func main() {

     str := strconv.FormatFloat(12342323.234232, 'f', 6, 64)

     fmt.Println(str)
 }

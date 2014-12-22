package main

 import (
     "fmt"
     "os"
 )

 func main() {
     file := "file.txt"

     if _, err := os.Stat(file); err == nil {
      fmt.Println(file, "exist!")
     } else {
      fmt.Println("No existe!")
     }

 }

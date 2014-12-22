package main

 import (
     "bitbucket.org/kardianos/osext"
     "fmt"
 )


 func main() {
     // get the current folder of the running program

     path, err := osext.ExecutableFolder()

     if err != nil {
        fmt.Println(err)
     }

     fmt.Println("Program is executing at folder :", path)
 }

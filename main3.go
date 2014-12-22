package main

import (
	"log"
	"fmt"
	"errors"
)

var ErrNoSearchRoot = errors.New("client: server doesn't support search")

func printInConsole(message string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(message)
	}
}

func getCurrentSpeed() (int, error) {
	//return 100, fmt.Errorf("this is a custom error")
	return 100, ErrNoSearchRoot
}

func main() {
	currentSpeed, err := getCurrentSpeed()
	if err != nil {
		log.Fatalf("could not get a client to fetch extra config: %v", err)
	}
	fmt.Println(currentSpeed)

}

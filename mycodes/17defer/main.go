package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	mydiffer()

}

func mydiffer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

}

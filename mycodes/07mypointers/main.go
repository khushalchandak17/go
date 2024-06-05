package main

import "fmt"

func main() {

	fmt.Println("Welcome to world of pointers")

	//  var one int
	// var ptr *int
	// var ptrs *string
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	var ptr = &myNumber

	// Refrence means &... simple
	// Note in "var ptr *int" we created only pointer, but above we are creating and also refrenceing it at the same time

	fmt.Println("Value of actual pointer is ", ptr)

	fmt.Println("Value in side pointer is ", *ptr)
	//fmt.Println("Value in side pointer in other way only to understand ", *&myNumber)

	// If you need the value inside pointer address simple add * in it

	*ptr = *ptr + 2
	fmt.Println("New value is ", myNumber)
}

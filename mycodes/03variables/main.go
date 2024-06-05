package main

import "fmt"

const Logtoken string = "zxccxzxzczxc" // Public

func main() {
	var username string = "Khushal"
	fmt.Println(username)
	fmt.Printf("Variable is of TYPE: %T \n", username)

	var isLoggedIN bool = true
	fmt.Println(isLoggedIN)
	fmt.Printf("Variable is of TYPE: %T \n", isLoggedIN)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of TYPE: %T \n", smallval)

	var smallFloat float32 = 255.4555455555
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of TYPE: %T \n", smallFloat)

	var smallFloat64 float64 = 255.4555455555
	fmt.Println(smallFloat64)
	fmt.Printf("Variable is of TYPE: %T \n", smallFloat64)

	// Default values and some aliases
	/*
		var anothervariables init
		fmt.Println(anothervariables)
		fmt.Printf("Variable is of TYPE: %T \n", anothervariables)
	*/

	// Implicit type
	var website = "google.com"
	fmt.Println(website)

	// no var style
	numberofuser := 3000000
	fmt.Println(numberofuser)

	fmt.Println(Logtoken)
}

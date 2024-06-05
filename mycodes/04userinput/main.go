package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	// Notes
	//comma okay syntax
	// comma error syntax
	// pgk.go.dev
	// https://pkg.go.dev/bufio#pkg-functions
	// https://pkg.go.dev/os#Chown

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma OK || comma err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating: ", input)
	fmt.Printf("Thanks for rat of type which is  %T", input)
}

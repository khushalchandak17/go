package main

import "fmt"

func main() {

	FizzBuzz(16)
}

func FizzBuzz(n int) {
	for i := 1; i < n; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Print("fizz buzz, ")
		} else if i%3 == 0 {
			fmt.Print("fizz, ")
		} else if i%5 == 0 {
			fmt.Print("buzz, ")
		} else {
			fmt.Print(i, ", ")
		}
	}

	if n%3 == 0 {
		fmt.Print("fizz")
	} else if n%5 == 0 {
		fmt.Print("buzz")
	} else {
		fmt.Print(n)
	}
	fmt.Println()
}

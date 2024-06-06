package main

import "fmt"

func main() {
	fmt.Println("Welcome to the world of functions")
	greeter()

	result := adder(3, 5)
	fmt.Println("The result is", result)

	proRes := proadder(2, 3, 4, 5, 6)
	fmt.Println("The pro result is ", proRes)

	proRes2, _ := proadder2(2, 3, 4, 5, 6)
	fmt.Println("The pro result2 is ", proRes2)

	proRes3, msg := proadder2(2, 3, 4, 5, 6)
	fmt.Println("The pro msg2 is ", msg, proRes3)
}

func adder(valone int, valtwo int) int {
	return valone + valtwo
}
func greeter() {
	fmt.Println("Namastey from Go")
}

func proadder(values ...int) int {

	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func proadder2(values ...int) (int, string) {

	total := 0
	for _, val := range values {
		total += val
	}
	return total, "This is a pro result by func"
}

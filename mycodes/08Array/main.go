package main

import "fmt"

func main() {
	fmt.Println("Welcome to world of array")

	var furitList [4]string
	furitList[0] = "Apple"
	furitList[1] = "Tomato"
	furitList[3] = "Peach"

	fmt.Println("Fruit list is ", furitList)
	fmt.Println("Fruit list is ", len(furitList))

	var vegList = [3]string{"Potato", "beans", "mushroom"}
	fmt.Println("Fruit list is ", vegList)

}

package main

import "fmt"

func main() {

	fmt.Println("Lets check num in list")

	var A = NumInList([]int{1, 2, 3, 4, 5}, 5)    //true
	var B = NumInList([]int{3, 3, 3, 3}, 5)       //false
	var C = NumInList([]int{1, 2, 3, 4, 6}, 5)    //false
	var D = NumInList([]int{1, 2, 5, 3, 4, 5}, 5) //true

	fmt.Println(A, B, C, D)
}

func NumInList(list []int, num int) bool {
	for _, val := range list {
		if val == 5 {
			return true
		}
	}
	return false
}

package main

import "fmt"

func main() {

	fmt.Println("Lets check num in list")

	var A = SumInList([]int{1, 2, 3, 4, 5})     //true
	var B = SumInList([]int{3, 3, 3, 3})        //false
	var C = SumInList([]int{1, 2, 3, 4, 6})     //false
	var D = Sum2InList([]int{1, 2, 5, 3, 4, 5}) //true

	fmt.Println(A, B, C, D)
}

func SumInList(list []int) int {
	x := 0
	for _, val := range list {
		x = val + x
		//x += val
	}
	return x

}

func Sum2InList(list []int) int {

	if len(list) == 0 {
		return 0
	}
	return list[0] + Sum2InList(list[1:])

}

package main

import "fmt"

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"

func main() {
	fmt.Println(Reverse("dog"))
	Reverse("cat")
}
func Reverse(word string) string {
	var res string
	// for _, r := range word {
	// 	res = string(r) + res
	// }

	for i := len(word) - 1; i >= 0; i-- {

		res = res + string(word[i])
	}
	return res
}

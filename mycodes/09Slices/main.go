package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to slices")

	var fruilLists = []string{"Apple", "Tomato", "Peach"}

	fmt.Printf("Type of frutilist is %T\n", fruilLists)

	fruilLists = append(fruilLists, "Mango", "Banana")
	fmt.Println(fruilLists)

	fruilLists = append(fruilLists[1:])
	fmt.Println(fruilLists)

	fruilLists = append(fruilLists[1:3])
	fmt.Println(fruilLists)
	// Note last -1 is only considerd

	highscores := make([]int, 4)

	highscores[0] = 111
	highscores[1] = 222
	highscores[2] = 212
	highscores[3] = 333

	fmt.Println(highscores)

	highscores = append(highscores, 444, 666)

	fmt.Println(highscores)

	sort.Ints(highscores)

	fmt.Println(highscores)

	fmt.Println(sort.IntsAreSorted(highscores))

	highscores = append(highscores, 101)

	fmt.Println(highscores)

	fmt.Println(sort.IntsAreSorted(highscores))

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Fprintln(os.Stdout, []any{"Welcome to our Pizza App"}...)

	fmt.Println("Please rate btw 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)

	// numRating, err := strconv.ParseFloat(input, 64)
	// correct way is
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The new rating is added by one is ", numRating+1)
	}

}

package main

import "fmt"

func main() {
	fmt.Println("welcome to loop")

	days := []string{"Sunday", "Mon", "Tue", "Wed", "Thu", "fri"}
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(i)
		fmt.Println(days[i])
	}

	for index, days := range days {
		fmt.Printf("Value for index %v is %v\n", index, days)
	}

	for _, days := range days {
		fmt.Printf("Value  is %v\n", days)
	}

	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 5 {
			break
		}

		fmt.Println("RGValue is ", rougueValue)
		rougueValue++
	}
	rougueValue2 := 1

	for rougueValue2 < 10 {
		if rougueValue2 == 2 {
			goto lco
		}

		fmt.Println("RGValue2 is ", rougueValue2)
		rougueValue2++
	}

lco:
	fmt.Println("Juming at learn code online using goto syntax")
	// You can also use continue instread of break
}

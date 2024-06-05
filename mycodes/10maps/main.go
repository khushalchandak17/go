package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"

	languages["RB"] = "Ruby"

	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println("Js stand for", languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	// Note this delete can also be used for slices

	// Loops are intresting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// comma ok syntax

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}
}

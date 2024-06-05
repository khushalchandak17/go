package main

import "fmt"

func main() {

	fmt.Println("Welcome to struct")

	// No inheritance in golang; No super or parent

	hitesh := User{"Hitesth", "asd@go.dev", true, 16}
	fmt.Println(hitesh)

	fmt.Printf("Hitest details are: %+v\n", hitesh)

	fmt.Printf("Hitest details are: %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v", hitesh.Name, hitesh.Email)

}

// Note first capital letters

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

package main

import "fmt"

func main() {

	fmt.Println("Welcome to struct")

	// No inheritance in golang; No super or parent

	hitesh := User{"Hitesth", "asd@go.dev", true, 16}
	fmt.Println(hitesh)

	fmt.Printf("Hitest details are: %+v\n", hitesh)

	fmt.Printf("Hitest details are: %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v\n", hitesh.Name, hitesh.Email)

	hitesh.getstatus()
	hitesh.newemail()

	fmt.Printf("Name is %v and old email is %v\n", hitesh.Name, hitesh.Email)

}

// Note first capital letters

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getstatus() {
	fmt.Println("The status for user is ", u.Status)

}

func (u User) newemail() {
	u.Email = "test@go.dev"
	fmt.Println("The New email for user is ", u.Email)

}

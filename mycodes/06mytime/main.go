package main

import (
	"fmt"
	"time"
)

func main() {

	// https://pkg.go.dev/time#example-Time.Format

	fmt.Println("Welcome to time study of go lang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// Note this is the fixed format

	fmt.Println(presentTime.Format("01-02-2006"))

	// Note this is the fixed format
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	//https://go.dev/src/time/format.go
	// https://pkg.go.dev/time#Time.Format

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Print(createdDate.Format("01-02-2006 Monday"))

}

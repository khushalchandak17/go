package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.google.com:443/learn?coursename=reactjs&paymentid=asdada12"

func main() {
	fmt.Println("Hi ")
	fmt.Println(myurl)

	//Parsing the URL
	result, _ := url.Parse(myurl)

	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port()) // Port is not a property its a method
	fmt.Println(result.RawQuery)

	qparams := result.Query() // special way to store query params

	fmt.Printf("The type of qparams are : %T\n", qparams)

	// So you see they are in form key:value

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is :", val)
	}

	//////// Lets construct url now
	// Note & is mandatroy
	partsOfurl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	// Note string is used like a method or what?
	anotherUrl := partsOfurl.String()
	fmt.Println(anotherUrl)
}

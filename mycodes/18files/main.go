package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Hi")
	content := "This needs to go in file - LCO"

	file, err := os.Create("./mylcogofile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)

	}
	fmt.Println("Length is", length)

	defer file.Close()
	readFile("./mylcogofile.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilerror(err)
	
	fmt.Println("Text data inside the file is \n", databyte)
	fmt.Println("Text data inside the file is \n", string(databyte))

}

/// Note using err condition is bit hard, better to create a function for it

func checkNilerror(err error) {
	if err != nil {
		panic(err)
	}
}

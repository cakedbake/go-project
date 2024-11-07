package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Specify the filename
	filename := "example.txt"

	// Read the file into a string
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the contents of the string
	fmt.Println(string(data))
}

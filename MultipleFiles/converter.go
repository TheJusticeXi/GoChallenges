package main

import (
	"fmt"
	"strings"
)

func ReadInput() string {
	var input string = ""
	fmt.Println("Please enter your input string:")
	fmt.Scanf("%s", &input)
	return input
}

func ConvertUpperCase(input string) string {

	var output string = strings.ToUpper(input)

	return output
}

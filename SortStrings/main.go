package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	fmt.Println("Please enter set of word:")
	var inputStr string
	fmt.Scanf("%s", &inputStr)

	tokens := strings.Split(inputStr, ",")

	sort.Strings(tokens)

	for i := range tokens {
		fmt.Printf("%s", tokens[i])

		if i != len(tokens)-1 {
			fmt.Printf(",")
		}
	}

	return
}

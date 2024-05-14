package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	fmt.Println("Please enter the input:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	tokens := strings.Split(input, " ")

	tokenMap := make(map[string]int)

	var output []string

	for i := range tokens {
		if tokenMap[tokens[i]] != 1 {
			tokenMap[tokens[i]] = 1
			output = append(output, tokens[i])

		}
	}

	sort.Strings(output)

	fmt.Println(strings.Join(output, " "))
}

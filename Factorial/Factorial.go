package main

import (
	"fmt"
)

func main() {

	fmt.Println("Please enter factorial number")
	var facNum int = 0
	fmt.Scanf("%d", &facNum)

	var facResult int = 0
	if facNum == 0 {
		facResult = 1
	} else if facNum > 0 {
		facResult = 1
		for index := 1; index <= facNum; index++ {
			facResult *= index
		}
	}

	fmt.Printf("Factorial Result = %d \n", facResult)

	return
}
